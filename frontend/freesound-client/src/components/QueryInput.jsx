import { useState } from 'react'

// Receives onSearch function as a prop from App.jsx
export default function QueryInput({ onSearch }) {
  const [searchId, setSearchId] = useState('')

  // This function runs when form is submitted
  function handleSubmit(e) {
    e.preventDefault() // Prevent the browser's default form behavior (page reload)

    // Call the onSearch function passed from App.jsx
    // This "lifts" the search action up to the parent component
    if (searchId.trim()) {
      onSearch(searchId)
    }
  }

  return (
    // onSubmit connects to our handleSubmit function
    <form onSubmit={handleSubmit} className="flex flex-row items-center gap-4 p-4 w-full max-w-md">
      <input
        type="text"
        placeholder="Enter single ID"
        className="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
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
