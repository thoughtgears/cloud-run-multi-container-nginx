import { useState } from 'react'

function App() {
  interface User {
    id: string
    first_name: string
    last_name: string
    username: string
    email: string
    created_at: string
    updated_at: string
  }

  const [pingResult, setPingResult] = useState<string | null>(null)
  const [users, setUsers] = useState<User[]>([])
  const [loadingUsers, setLoadingUsers] = useState<boolean>(false)
  const [usersError, setUsersError] = useState<string | null>(null)

  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || ''
  console.log('apiBaseUrl', apiBaseUrl)

  const handlePing = async () => {
    try {
      const response = await fetch(`${apiBaseUrl}/ping`)
      if (response.ok) {
        const data = await response.text()
        setPingResult(`Response: ${data}`)
      } else {
        setPingResult(`Error: ${response.status} - ${response.statusText}`)
      }
    } catch (error) {
      setPingResult(`Fetch Error: ${error}`)
    }
  }

  const handleGetUsers = async () => {
    setLoadingUsers(true)
    setUsersError(null)
    try {
      const response = await fetch(`${apiBaseUrl}/users`)
      if (response.ok) {
        const data = await response.json()
        setUsers(data)
      } else {
        setUsersError(`Error: ${response.status} - ${response.statusText}`)
      }
    } catch (error) {
      setUsersError(`Fetch Error: ${error}`)
    } finally {
      setLoadingUsers(false)
    }
  }

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">
        Test the NGINX Reverse Proxy with Go API
      </h1>

      <div className="mb-6 border p-4">
        <h2 className="text-xl font-semibold mb-2">Ping API</h2>
        <button
          onClick={handlePing}
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          Ping the Server
        </button>
        <div id="ping-result" className="mt-2">
          {pingResult}
        </div>
      </div>

      <div className="border p-4">
        <h2 className="text-xl font-semibold mb-2">Users API</h2>
        <button
          onClick={handleGetUsers}
          className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mb-2"
          disabled={loadingUsers}
        >
          {loadingUsers ? 'Loading Users...' : 'Get All Users'}
        </button>
        {usersError && <div className="text-red-500">{usersError}</div>}
        <div>
          {users.length > 0 ? (
            <ul className="list-disc pl-5">
              {users.map((user) => (
                <li key={user.id}>
                  ID: <span className="font-semibold">{user.id}</span>, First
                  Name: <span className="font-semibold">{user.first_name}</span>
                  , Last Name:{' '}
                  <span className="font-semibold">{user.last_name}</span>,
                  Username:{' '}
                  <span className="font-semibold">{user.username}</span>, Email:{' '}
                  <span className="font-semibold">{user.email}</span>, Created
                  At: <span className="text-gray-500">{user.created_at}</span>,
                  Updated At:{' '}
                  <span className="text-gray-500">{user.updated_at}</span>
                </li>
              ))}
            </ul>
          ) : (
            !loadingUsers && <div>No users found.</div>
          )}
        </div>
      </div>
    </div>
  )
}

export default App
