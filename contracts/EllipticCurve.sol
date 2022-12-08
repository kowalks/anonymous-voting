// SPDX-License-Identifier: MIT

pragma solidity >=0.8;

library EllipticCurve {
    // EllipticCurve y^2 = x^3 + a x^2 + b
    // Over Galois Field GF(p)
    // Name: secp256k1
    uint public constant A = 70;
    uint public constant B = 0;
    uint public constant P = 71;

    // invMod by Fermat Little Theorem: a^(p-2) is the inverse of a mod p
    function modInv(uint a) internal pure returns (uint) {
        return modExp(a, P - 2);
    }

    // computes a^b mod p in O(log b)
    function modExp(uint a, uint b) internal pure returns (uint) {
        if (b == 0) return 1;

        uint c = modExp(a, b / 2);
        c = (c * c) % P;
        if (b % 2 != 0) c = (c * a) % P;
        return c;
    }

    function modMul(uint a, uint b) internal pure returns (uint) {
        return (a * b) % P;
    }

    function modDiv(uint a, uint b) internal pure returns (uint) {
        return modMul(a, modInv(b));
    }

    function modAdd(uint a, uint b) internal pure returns (uint) {
        return (a + b) % P;
    }

    function modSub(uint a, uint b) internal pure returns (uint) {
        return (a + P - b) % P;
    }

    function pAdd(
        uint xp,
        uint yp,
        uint xq,
        uint yq
    ) internal pure returns (uint, uint) {
        if (xp == xq) {
            return pDbl(xp, yp);
        }

        uint lambda = modDiv(modSub(yq,yp), modSub(xq,xp));
        uint xr = modSub(modExp(lambda, 2), modAdd(xp, xq));
        uint yr = modSub(modMul(lambda, modSub(xp, xr)),yp);
        return (xr, yr);
    }

    function pDbl(uint xp, uint yp)
        internal
        pure
        returns (uint, uint)
    {
        uint xq = xp;

        uint lambda = modDiv(3 * modExp(xp, 2) + A, modMul(yp, 2));
        uint xr = modSub(modExp(lambda, 2), modAdd(xp, xq));
        uint yr = modMul(lambda, modSub(xp, xr)) - yp;
        return (xr, yr);
    }

    function pNeg(uint x, uint y)
        internal
        pure
        returns (uint, uint)
    {
        return (x, modSub(0, y));
    }

    function pMul(
        uint x,
        uint y,
        uint k
    ) internal pure returns (uint, uint) {
        if (k == 1) return (x, y);

        uint xx;
        uint yy;
        (xx, yy) = pMul(x, y, k / 2);
        (xx, yy) = pDbl(xx, yy);

        if (k % 2 != 0) (xx, yy) = pAdd(xx, yy, x, y);
        return (xx, yy);
    }
}
