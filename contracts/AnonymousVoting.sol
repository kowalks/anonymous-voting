// SPDX-License-Identifier: MIT

pragma solidity >=0.8;

struct Point {
    uint x;
    uint y;
}

struct Candidate {
    string name;
    Point pubkey;
}

// Voter model: ea
struct Voter {
    Point pubkey;
}

// Key Manager model
struct Manager {
    address payable owner;
    Point pubkey;
    Point announcedKey;
    uint privateKey;
}

// Vote model
struct Ballot {
    Point SA;
    Point R;
    string signature;
}

enum Phase {
    Setup,
    Vote,
    Tally
}

contract Curve {
    uint P = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF000000000000000000000001;
    uint N = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFF16A2E0B8F03E13DD29455C5C2A3D;
    uint A = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFE;
    uint B = 0xB4050A850C04B3ABF54132565044B0B7D7BFD8BA270B39432355FFB4;
    uint Gx = 0xB70E0CBD6BB4BF7F321390B94A03C1D356C21122343280D6115C1D21;
    uint Gy = 0xBD376388B5F723FB4C22DFE6CD4375A05A07476444D5819985007E34;
    uint BitSize = 224;
    string name = "P-224";
}

contract AnonymousVoting {
    Phase public phase = Phase.Setup;

    mapping(uint => Candidate) public candidates;
    mapping(uint => Voter) public voters;
    mapping(uint => Ballot) public ballots;
    mapping(uint => Manager) public managers;

    uint public candidateCount = 0;
    uint public voterCount = 0;
    uint public ballotCount = 0;
    uint public managerCount = 0;

    function addCandidate (string memory name, uint x, uint y) public {
        require(phase == Phase.Setup);

        candidates[candidateCount] = Candidate(name, Point(x, y));
        candidateCount++;
    }

    function addVoter (uint x, uint y) public {
        require(phase == Phase.Setup);

        voters[voterCount] = Voter(Point(x, y));
        voterCount++;
    }

    // Two key managers announce rG publicly
    function registerManager(uint x, uint y) public payable {
        require(managerCount < 2, "Too many managers");
        require(phase == Phase.Setup);

        address payable owner = payable(msg.sender);
        managers[managerCount] = Manager(owner, Point(x, y), Point(0, 0), 0);
        managerCount++;
    }

    function announcePublicKey(uint x, uint y) public {
        require(phase == Phase.Setup);

        address payable owner = payable(msg.sender);
        
        for (uint i=0; i < managerCount; i++) {
            if (managers[i].owner == owner) {
                managers[i].announcedKey.x = x;
                managers[i].announcedKey.x = y;
            }
        }
    }

    function getPublicKey() public returns (uint, uint) {
        uint x = managers[0].announcedKey.x;
        uint y = managers[0].announcedKey.y;
        for (uint i=0; i < managerCount; i++) {
            require(managers[i].announcedKey.x != x, "divergent announcement");
            require(managers[i].announcedKey.y != y, "divergent announcement");
        }

        phase = Phase.Vote;
        return (x,y);
    }

    function vote(uint SAx, uint SAy, uint Rx, uint Ry, string memory signature) public {
        require(ballotCount < voterCount, "too many votes");
        require(phase == Phase.Vote);

        Point memory SA = Point(SAx, SAy);
        Point memory R = Point(Rx, Ry);

        ballots[ballotCount] = Ballot(SA, R, signature);
        ballotCount++;
    }

    function revealPublicKey(uint r) public {
        require(ballotCount == voterCount, "missing ballots");
        
        phase = Phase.Tally;

        address payable owner = payable(msg.sender);

        for (uint i=0; i < managerCount; i++) {
            if (managers[i].owner == owner) {
                managers[i].privateKey = r;
            }
        }
    }

} 