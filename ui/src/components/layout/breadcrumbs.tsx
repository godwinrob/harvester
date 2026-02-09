import { Link, useLocation } from '@tanstack/react-router'
import { ChevronRight, Home } from 'lucide-react'

// Map routes to breadcrumb labels
const routeLabels: Record<string, string> = {
  '/': 'Dashboard',
  '/users': 'Users',
  '/galaxies': 'Galaxies',
  '/resources': 'Resources',
}

export function Breadcrumbs() {
  const location = useLocation()
  const pathSegments = location.pathname.split('/').filter(Boolean)

  // Build breadcrumb items
  const items = pathSegments.map((segment, index) => {
    const path = '/' + pathSegments.slice(0, index + 1).join('/')
    const label = routeLabels[path] || segment

    return {
      path,
      label: label.charAt(0).toUpperCase() + label.slice(1),
      isLast: index === pathSegments.length - 1,
    }
  })

  // Add home if not on home page
  if (location.pathname !== '/') {
    items.unshift({ path: '/', label: 'Dashboard', isLast: false })
  }

  if (items.length === 0) {
    return null
  }

  return (
    <nav aria-label="Breadcrumb" className="flex items-center text-sm">
      {items.map((item, index) => (
        <div key={item.path} className="flex items-center">
          {index > 0 && (
            <ChevronRight className="mx-2 h-4 w-4 text-muted-foreground" />
          )}
          {item.isLast ? (
            <span className="font-medium text-foreground">{item.label}</span>
          ) : (
            <Link
              to={item.path}
              className="text-muted-foreground hover:text-foreground transition-colors"
            >
              {index === 0 && items.length > 1 ? (
                <span className="flex items-center gap-1">
                  <Home className="h-4 w-4" />
                  <span className="sr-only sm:not-sr-only">{item.label}</span>
                </span>
              ) : (
                item.label
              )}
            </Link>
          )}
        </div>
      ))}
    </nav>
  )
}
