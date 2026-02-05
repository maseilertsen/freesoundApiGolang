import { useState } from 'react'

export default function ResultsDisplay({ results }) {
  const [copied, setCopied] = useState(false)

  if (!results) {
    return null
  }

  // Handle both single (name/username) and multi (text) result formats
  const displayText = results.text
    ? results.text
    : `${results.name} - ${results.username}`

  function handleCopy() {
    navigator.clipboard.writeText(displayText).then(() => {
      setCopied(true)
      setTimeout(() => setCopied(false), 2000)
    })
  }

  return (
    <div className="w-full max-w-2xl mx-auto p-4">
      <div className="flex justify-between items-center mb-2">
        <h2 className="text-lg font-semibold text-gray-700">Result</h2>
        <button
          onClick={handleCopy}
          className="px-3 py-1 text-sm bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition-colors"
        >
          {copied ? 'Copied!' : 'Copy'}
        </button>
      </div>

      {/* Code block styling with plain text content */}
      <div className="bg-gray-100 p-4 rounded-lg select-all">
        <p className="text-gray-800 font-mono">
          {displayText}
        </p>
      </div>
    </div>
  )
}
