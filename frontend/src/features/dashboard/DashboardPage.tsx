import { useNavigate } from 'react-router-dom';
import { useAuthStore } from '../../stores/authStore';

export const DashboardPage = () => {
  const { user, logout } = useAuthStore();
  const navigate = useNavigate();

  const handleLogout = async () => {
    await logout();
    navigate('/login');
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b-2 border-gray-200">
        <div className="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
          <h1 className="text-2xl font-bold text-gray-900">DadMail</h1>
          <div className="flex items-center gap-4">
            <span className="text-lg text-gray-700">
              {user?.full_name}
            </span>
            <button onClick={handleLogout} className="btn-secondary">
              Logout
            </button>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-6 py-8">
        <div className="card">
          <h2 className="text-3xl font-bold text-gray-900 mb-4">
            Welcome, {user?.full_name}!
          </h2>
          <p className="text-xl text-gray-700 mb-6">
            Your inbox is ready.
          </p>

          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3 mt-8">
            {/* Categories */}
            {[
              { name: 'Medical', color: 'bg-senior-medical', icon: '💊' },
              { name: 'Financial', color: 'bg-senior-financial', icon: '💰' },
              { name: 'Family', color: 'bg-senior-family', icon: '👨‍👩‍👧‍👦' },
              { name: 'Commercial', color: 'bg-senior-commercial', icon: '🛒' },
              { name: 'Administrative', color: 'bg-senior-admin', icon: '📋' },
              { name: 'Spam', color: 'bg-senior-spam', icon: '🗑️' },
            ].map((category) => (
              <div
                key={category.name}
                className="card hover:shadow-lg transition-shadow cursor-pointer"
              >
                <div className={`w-16 h-16 ${category.color} rounded-senior flex items-center justify-center text-3xl mb-4`}>
                  {category.icon}
                </div>
                <h3 className="text-xl font-bold text-gray-900 mb-2">
                  {category.name}
                </h3>
                <p className="text-base text-gray-600">0 messages</p>
              </div>
            ))}
          </div>

          <div className="mt-12 p-6 bg-blue-50 rounded-senior border-2 border-blue-200">
            <p className="text-lg text-gray-700">
              <strong>Coming Soon:</strong> Email integration with Gmail and IMAP/SMTP providers.
              Your inbox will automatically categorize emails to reduce cognitive overhead.
            </p>
          </div>
        </div>
      </main>
    </div>
  );
};
