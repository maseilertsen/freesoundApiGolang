import { useState } from 'react'
import './App.css'
import Header from './components/Header.jsx'
import QueryInput from './components/QueryInput.jsx'
import ResultsDisplay from './components/ResultsDisplay.jsx'

function App() {
  // State to hold API results - lives here so both components can access it
  const [results, setResults] = useState(null)

  // This function will be passed to QueryInput
  // TODO: make it call the actual API
  function handleSearch(searchId) {
    console.log('App received search for:', searchId)

    // For now, set mock data to test ResultsDisplay
    // TODO: Replace this with actual API call
    setResults({
      id: searchId,
      name: 'Test Sound',
      username: 'test_user',
      message: 'This is mock data - API connection coming next!'
    })
  }

  return (
    <div className="min-h-screen bg-white flex flex-col items-center pt-32">
      <Header />
      {/* Pass handleSearch function down to QueryInput */}
      <QueryInput onSearch={handleSearch} />
      {/* Pass results down to ResultsDisplay */}
      <ResultsDisplay results={results} />
    </div>
  )
}
export default App
