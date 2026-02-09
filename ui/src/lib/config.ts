/**
 * Application configuration with validation
 * Fails fast if critical environment variables are misconfigured
 */

interface Config {
  apiUrl: string
  appName: string
  appVersion: string
}

/**
 * Validates environment variable and returns its value
 * Throws error if validation fails
 */
function validateEnvVar(
  name: string,
  value: string | undefined,
  options: {
    required?: boolean
    defaultValue?: string
    invalidPatterns?: string[]
  } = {}
): string {
  const { required = false, defaultValue, invalidPatterns = [] } = options

  // Use default if provided and value is undefined
  if (value === undefined || value === '') {
    if (defaultValue !== undefined) {
      return defaultValue
    }
    if (required) {
      throw new Error(createConfigError(name, 'is required but not set'))
    }
    return ''
  }

  // Check for placeholder patterns
  const placeholderPatterns = [
    'CHANGE_ME',
    'TODO',
    'PLACEHOLDER',
    'your_',
    'example.com',
    ...invalidPatterns,
  ]

  for (const pattern of placeholderPatterns) {
    if (value.includes(pattern)) {
      throw new Error(
        createConfigError(
          name,
          `contains placeholder value "${pattern}" - please update with actual configuration`
        )
      )
    }
  }

  return value
}

/**
 * Creates a formatted configuration error message
 */
function createConfigError(varName: string, reason: string): string {
  return `
┌─────────────────────────────────────────────────────────────────┐
│ CONFIGURATION ERROR                                             │
├─────────────────────────────────────────────────────────────────┤
│ Environment variable: ${varName.padEnd(42)} │
│ Issue: ${reason.padEnd(51)} │
│                                                                 │
│ To fix:                                                         │
│   1. Copy ui/.env.example to ui/.env                            │
│   2. Update the ${varName} value                     │
│   3. Restart the development server                             │
│                                                                 │
│ For local development:                                          │
│   VITE_API_URL should be /v1 (uses vite proxy)                  │
│                                                                 │
│ For production:                                                 │
│   VITE_API_URL should be your API server URL                    │
└─────────────────────────────────────────────────────────────────┘
`
}

/**
 * Load and validate configuration
 * Call this early in app initialization to fail fast
 */
export function loadConfig(): Config {
  try {
    const config: Config = {
      apiUrl: validateEnvVar('VITE_API_URL', import.meta.env.VITE_API_URL, {
        defaultValue: '/v1',
      }),
      appName: validateEnvVar('VITE_APP_NAME', import.meta.env.VITE_APP_NAME, {
        defaultValue: 'Harvester',
      }),
      appVersion: validateEnvVar(
        'VITE_APP_VERSION',
        import.meta.env.VITE_APP_VERSION,
        {
          defaultValue: 'dev',
        }
      ),
    }

    // Log configuration in development
    if (import.meta.env.DEV) {
      console.log('✅ Configuration loaded successfully:', {
        apiUrl: config.apiUrl,
        appName: config.appName,
        appVersion: config.appVersion,
      })
    }

    return config
  } catch (error) {
    // Log to console for visibility
    console.error(error)

    // Re-throw to stop application initialization
    throw error
  }
}

// Export singleton config instance
export const config = loadConfig()
