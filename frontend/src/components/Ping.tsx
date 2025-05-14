// src/components/PingSection.tsx
import { useState } from 'react'
import { useAuth } from '../contexts/AuthContext' // Import useAuth

const Ping = () => {
  const [pingResult, setPingResult] = useState<string | null>(null)
  const [pingLoading, setPingLoading] = useState<boolean>(false)
  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || ''
  const { currentUser } = useAuth() // Get the current Firebase user

  const handlePing = async () => {
    if (!currentUser) {
      setPingResult(
        'Error: User not logged in. Cannot ping authenticated endpoint.'
      )
      return
    }

    setPingLoading(true)
    setPingResult(null) // Clear previous result

    try {
      const idToken = await currentUser.getIdToken() // Get the Firebase ID token

      const response = await fetch(`${apiBaseUrl}/ping`, {
        headers: {
          Authorization: `Bearer ${idToken}`, // Add the Authorization header
        },
      })

      const responseData = await response.json() // Assuming the response is JSON now

      if (response.ok) {
        // If your backend sends back JSON like { message: 'pong (authenticated)', uid: '...' }
        setPingResult(
          `Response: ${responseData.message}, UID: ${responseData.uid || 'N/A'}`
        )
      } else {
        // Handle errors from the backend (e.g., 401, 403)
        setPingResult(
          `Error: ${response.status} - ${responseData.message || response.statusText}`
        )
      }
    } catch (error: any) {
      console.error('Fetch Error:', error)
      setPingResult(`Workspace Error: ${error.message || error}`)
    } finally {
      setPingLoading(false)
    }
  }

  return (
    <div className="mb-6 border p-4">
      <h2 className="text-xl font-semibold mb-2">Ping API (Authenticated)</h2>
      <button
        onClick={handlePing}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        disabled={pingLoading || !currentUser} // Disable if loading or no user
      >
        {pingLoading
          ? 'Pinging...'
          : currentUser
            ? 'Ping the Server'
            : 'Login to Ping'}
      </button>
      <div id="ping-result" className="mt-2">
        {pingResult}
      </div>
      {!currentUser && (
        <p className="text-sm text-yellow-600 mt-2">
          You must be logged in to ping the server.
        </p>
      )}
    </div>
  )
}

export default Ping
