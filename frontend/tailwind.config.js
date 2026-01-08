/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      // Senior-friendly font sizes
      fontSize: {
        'xs': ['0.875rem', { lineHeight: '1.5' }],
        'sm': ['1rem', { lineHeight: '1.5' }],
        'base': ['1.125rem', { lineHeight: '1.6' }],      // 18px - larger base
        'lg': ['1.25rem', { lineHeight: '1.6' }],         // 20px
        'xl': ['1.5rem', { lineHeight: '1.6' }],          // 24px
        '2xl': ['1.875rem', { lineHeight: '1.5' }],       // 30px
        '3xl': ['2.25rem', { lineHeight: '1.4' }],        // 36px
        '4xl': ['3rem', { lineHeight: '1.3' }],           // 48px
      },
      // High contrast colors for accessibility
      colors: {
        // Senior-friendly palette with high contrast
        'senior': {
          'primary': '#2563eb',      // Blue - easy to see
          'secondary': '#16a34a',    // Green - positive actions
          'danger': '#dc2626',       // Red - warnings/important
          'warning': '#ea580c',      // Orange - caution
          'info': '#0891b2',         // Cyan - information
          'medical': '#dc2626',      // Medical category
          'financial': '#16a34a',    // Financial category
          'family': '#2563eb',       // Family category
          'commercial': '#9333ea',   // Commercial category
          'admin': '#ea580c',        // Administrative category
          'spam': '#6b7280',         // Spam category
        },
      },
      // Larger spacing for easier interaction
      spacing: {
        '18': '4.5rem',
        '22': '5.5rem',
        '26': '6.5rem',
        '30': '7.5rem',
      },
      // Border radius - less rounded for clarity
      borderRadius: {
        'senior': '0.5rem',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}
