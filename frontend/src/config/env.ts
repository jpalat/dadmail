// Environment configuration

interface EnvConfig {
  apiUrl: string;
  environment: string;
  isDevelopment: boolean;
  isProduction: boolean;
}

const env: EnvConfig = {
  apiUrl: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1',
  environment: import.meta.env.MODE || 'development',
  isDevelopment: import.meta.env.DEV,
  isProduction: import.meta.env.PROD,
};

export default env;
