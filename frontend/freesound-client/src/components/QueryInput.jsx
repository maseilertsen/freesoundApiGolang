import { useState } from 'react'

export default function QueryInput() {
  const [searchId, setSearchId] = useState('')  

  // This function runs when form is submitted
  function handleSubmit(e) {
    e.preventDefault()// Prevent the browser's default form behavior (page reload)

    // Logging value before as proof of concept.
    console.log('Searching for ID:', searchId)
  }

  return (
    // onSubmit connects to our handleSubmit function
    <form onSubmit={handleSubmit} className="flex flex-col items-center gap-4 p-4">
      <input
        type="text"
        placeholder="Enter Freesound ID (e.g., 1234)"
        className="w-full max-w-md px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        value={searchId} // The input always displays what's in our state
        onChange={(e) => setSearchId(e.target.value)} // onChange fires every time the user types a character
      />
      <button
        type="submit"
        className="px-6 py-2 bg-blue-400 text-white font-semibold rounded-lg hover:bg-blue-650 transition-colors"
      >
        Search
      </button>
    </form>
  )
}
