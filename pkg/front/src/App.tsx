import { MainPage } from '@pages/main'
import React from 'react'
import { QueryClient, QueryClientProvider } from 'react-query'
import './App.css'

const queryClient = new QueryClient()

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        <MainPage />
      </div>
    </QueryClientProvider>
  )
}

export default App
