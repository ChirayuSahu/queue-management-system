
export default function LoginCard() {
  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="w-full max-w-5xl rounded-2xl border border-gray-200 shadow-sm overflow-hidden grid grid-cols-1 md:grid-cols-2">
        <div className="hidden md:block">
          <img
            src="https://images.unsplash.com/photo-1508780709619-79562169bc64"
            alt="Autumn forest"
            className="h-full w-full object-cover"
          />
        </div>
        <div className="flex items-center justify-center p-8 md:p-12">
          <div className="w-full max-w-md">
            <h1 className="text-3xl font-bold text-gray-900 mb-8">
              Login
            </h1>

            <form className="space-y-6">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Email
                </label>
                <input
                  type="email"
                  placeholder="Enter your email"
                  className="w-full rounded-lg border border-gray-300 px-4 py-3 focus:outline-none focus:ring-2 focus:ring-green-500"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Password
                </label>
                <input
                  type="password"
                  placeholder="Enter your password"
                  className="w-full rounded-lg border border-gray-300 px-4 py-3 focus:outline-none focus:ring-2 focus:ring-green-500"
                />
              </div>
              <button
                type="submit"
                className="w-full rounded-lg bg-green-500 py-3 text-white font-semibold hover:bg-green-700 transition"
              >
                Sign In
              </button>
            </form>
            <div className="block text-center mt-4 text-green-500">
              Don't have an account? <a href="/signup" className=" hover:text-green-700">Sign Up</a>
            </div>
            
          </div>
        </div>
      </div>
      
    </div>
    
  );
}
