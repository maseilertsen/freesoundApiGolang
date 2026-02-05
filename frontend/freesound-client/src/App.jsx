import { useState } from 'react'
import './App.css'
import Header from './components/Header.jsx'
import QueryInput from './components/QueryInput.jsx'
import ResultsDisplay from './components/ResultsDisplay.jsx'
import MultilineInput from './components/MultilineInput.jsx'

const API_BASE_URL = 'http://localhost:8080'

function App() {
  // State to hold API results - lives here so both components can access it
  const [results, setResults] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  // This function will be passed to QueryInput
  async function handleSearch(searchId) {
    console.log('App received search for:', searchId)

    // Reset states before new request
    setLoading(true)
    setError(null)
    setResults(null)

    // 
    try {
      const response = await fetch(API_BASE_URL + '/song/' + searchId + '/')
      console.log('Fetch url: '+ API_BASE_URL + '/song/' + searchId + '/')

        if (!response.ok) {
          throw new Error(`HTTP error! status:  ${response.status}`)
        }
        
        const result = await response.json()
        setResults(result)
      } catch (error) {
        setError(error.message)
      } finally{
        setLoading(false)
      }
  }

  if (loading){
    return <p>loading... data</p>
    }
    if (error){
      return <p>Error loading data</p>
    }


  return (
    <div className="min-h-screen bg-white flex flex-col items-center pt-32">
      <Header />
      {/* Pass handleSearch function down to QueryInput */}
      <QueryInput onSearch={handleSearch} />
      <MultilineInput  />
      {/* Pass results down to ResultsDisplay */}
      <ResultsDisplay results={results} />
    </div>
  )
}
export default App
