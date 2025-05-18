import admin from 'firebase-admin';

const firebaseAuthMiddleware = async (req, res, next) => {
    const authHeader = req.headers['x-firebase-authorization'];

    if (!authHeader || !authHeader.startsWith('Bearer ')) {
        return res.status(401).json({ message: 'Unauthorized: No token provided or incorrect format.' });
    }

    const idToken = authHeader.split('Bearer ')[1];

    if (!idToken) {
        return res.status(401).json({ message: 'Unauthorized: No token found after Bearer.' });
    }

    try {
        req.user = await admin.auth().verifyIdToken(idToken);
        next();
    } catch (error) {
        console.error('Error verifying Firebase ID token:', error.message);
        let errorMessage = 'Unauthorized: Invalid token.';
        if (error.code === 'auth/id-token-expired') {
            errorMessage = 'Unauthorized: Token expired.';
        }
        return res.status(403).json({ message: errorMessage, code: error.code });
    }
};

export default firebaseAuthMiddleware;