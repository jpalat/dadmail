function App() {
  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center p-8">
      <div className="card max-w-2xl w-full text-center">
        <h1 className="text-4xl font-bold text-gray-900 mb-6">
          DadMail
        </h1>
        <p className="text-xl text-gray-700 mb-8">
          Email designed for seniors
        </p>
        <div className="space-y-4">
          <p className="text-lg text-gray-600">
            Lower cognitive overhead • Clear labeling • Senior-friendly design
          </p>
          <div className="flex gap-4 justify-center mt-8">
            <button className="btn-primary">
              Get Started
            </button>
            <button className="btn-secondary">
              Learn More
            </button>
          </div>
        </div>
        <div className="mt-12 p-6 bg-blue-50 rounded-senior border-2 border-blue-200">
          <p className="text-base text-gray-700">
            <strong>Phase 1: Foundation</strong> - Backend and Frontend setup complete.
            Next: Email integration, categorization engine, and senior-friendly UI.
          </p>
        </div>
      </div>
    </div>
  )
}

export default App
