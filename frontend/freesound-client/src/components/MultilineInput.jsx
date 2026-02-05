import { useState } from 'react'

// Takes raw text input and returns an array of ID strings.
// Handles: newlines, commas, spaces as separators, and filters out empty strings.
function parseIds(rawText) {
  return rawText.split(/[\n,\s]+/).filter(id => id.length > 0)  
}

export default function MultilineInput({ onSearch }) {
  const [textAreaValue, setTextAreaValue] = useState('')

  function handleSubmit(e) {
    e.preventDefault()

    const ids = parseIds(textAreaValue)
    if (ids.length > 0) {
      onSearch(ids)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-4 p-4 w-full max-w-md">
      <textarea
        value={textAreaValue}
        onChange={(e) => setTextAreaValue(e.target.value)}
        className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        rows={5}
        placeholder="Enter IDs (separated by newlines, commas, or spaces)"
      />
      <button
        type="submit"
        className="self-end px-6 py-2 bg-blue-400 text-white font-semibold rounded-lg hover:bg-blue-650 transition-colors"
      >
        Search
      </button>
    </form>
  )
}

