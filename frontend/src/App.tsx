// src/App.tsx
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from 'react-router-dom'
import { AuthProvider, useAuth } from './contexts/AuthContext'
import HomePage from './pages/Home'
import LoginPage from './pages/Login'
import ProtectedRoute from './components/ProtectedRoute'

// A small component to handle redirection for already logged-in users trying to access /login
const LoginPageWrapper = () => {
  const { currentUser, loading } = useAuth()

  if (loading) {
    return <div>Loading...</div> // Or a spinner
  }

  return !currentUser ? <LoginPage /> : <Navigate to="/" replace />
}

function AppContent() {
  return (
    <Routes>
      <Route path="/login" element={<LoginPageWrapper />} />
      <Route element={<ProtectedRoute />}>
        <Route path="/" element={<HomePage />} />
        {/* Add other protected routes here if needed */}
      </Route>
      {/* Fallback for any other route */}
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}

function App() {
  return (
    <Router>
      <AuthProvider>
        <AppContent />
      </AuthProvider>
    </Router>
  )
}

export default App
