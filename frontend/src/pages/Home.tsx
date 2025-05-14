import Ping from '../components/Ping'
import Users from '../components/Users'
import { auth } from '../lib/firebase'
import { signOut } from 'firebase/auth'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../contexts/AuthContext'

const HomePage = () => {
  const navigate = useNavigate()
  const { currentUser } = useAuth()

  const handleLogout = async () => {
    try {
      await signOut(auth)
      navigate('/login')
    } catch (error) {
      console.error('Failed to logout:', error)
    }
  }

  return (
    <div className="container mx-auto p-4">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">
          Test the NGINX Reverse Proxy with Go API
        </h1>
        {currentUser && (
          <button
            onClick={handleLogout}
            className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
          >
            Logout ({currentUser.email})
          </button>
        )}
      </div>
      <Ping />
      <Users />
    </div>
  )
}

export default HomePage
