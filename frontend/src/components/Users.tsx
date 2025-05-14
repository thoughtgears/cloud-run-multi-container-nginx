import { useState } from 'react'
import type { User } from '../types'

const Users = () => {
  const [users, setUsers] = useState<User[]>([])
  const [loadingUsers, setLoadingUsers] = useState<boolean>(false)
  const [usersError, setUsersError] = useState<string | null>(null)
  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || ''

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
      setUsersError(`Workspace Error: ${error}`)
    } finally {
      setLoadingUsers(false)
    }
  }

  return (
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
                Name: <span className="font-semibold">{user.first_name}</span>,
                Last Name:{' '}
                <span className="font-semibold">{user.last_name}</span>,
                Username: <span className="font-semibold">{user.username}</span>
                , Email: <span className="font-semibold">{user.email}</span>,
                Created At:{' '}
                <span className="text-gray-500">{user.created_at}</span>,
                Updated At:{' '}
                <span className="text-gray-500">{user.updated_at}</span>
              </li>
            ))}
          </ul>
        ) : (
          !loadingUsers && !usersError && <div>No users found.</div>
        )}
      </div>
    </div>
  )
}

export default Users
