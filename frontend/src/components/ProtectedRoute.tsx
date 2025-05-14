import React from 'react'
import { Navigate, Outlet } from 'react-router-dom'
import { useAuth } from '../contexts/AuthContext'

const ProtectedRoute: React.FC = () => {
  const { currentUser, loading } = useAuth()

  if (loading) {
    // You can return a loading spinner here if you want
    return <div>Loading...</div>
  }

  return currentUser ? <Outlet /> : <Navigate to="/login" replace />
}

export default ProtectedRoute
