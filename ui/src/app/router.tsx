import {
  createRouter,
  createRoute,
  createRootRoute,
  Outlet,
  Link,
} from '@tanstack/react-router'
import { AppShell } from '@/components/layout/app-shell'
import { UserTable } from '@/features/users'
import { GalaxyTable } from '@/features/galaxies'
import { ResourceTable } from '@/features/resources'

// Root route with app shell layout
const rootRoute = createRootRoute({
  component: () => <Outlet />,
})

// Layout route for authenticated pages
const layoutRoute = createRoute({
  getParentRoute: () => rootRoute,
  id: 'layout',
  component: () => (
    <AppShell>
      <Outlet />
    </AppShell>
  ),
})

// Dashboard / Home
const indexRoute = createRoute({
  getParentRoute: () => layoutRoute,
  path: '/',
  component: () => (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
        <p className="text-muted-foreground">
          Welcome to Harvester. Manage your users, galaxies, and resources.
        </p>
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <DashboardCard
          title="Users"
          description="Manage user accounts and roles"
          href="/users"
          icon="users"
        />
        <DashboardCard
          title="Galaxies"
          description="Create and manage galaxies"
          href="/galaxies"
          icon="globe"
        />
        <DashboardCard
          title="Resources"
          description="Track and verify resources"
          href="/resources"
          icon="package"
        />
      </div>
    </div>
  ),
})

// Users routes
const usersRoute = createRoute({
  getParentRoute: () => layoutRoute,
  path: 'users',
  component: () => <Outlet />,
})

const usersIndexRoute = createRoute({
  getParentRoute: () => usersRoute,
  path: '/',
  component: () => (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Users</h1>
          <p className="text-muted-foreground">
            Manage user accounts and permissions
          </p>
        </div>
      </div>
      <UserTable />
    </div>
  ),
})

const userDetailRoute = createRoute({
  getParentRoute: () => usersRoute,
  path: '$userId',
  component: () => {
    const { userId } = userDetailRoute.useParams()
    return (
      <div>
        <h1 className="text-2xl font-bold">User Details</h1>
        <p className="text-muted-foreground">User ID: {userId}</p>
      </div>
    )
  },
})

// Galaxies routes
const galaxiesRoute = createRoute({
  getParentRoute: () => layoutRoute,
  path: 'galaxies',
  component: () => <Outlet />,
})

const galaxiesIndexRoute = createRoute({
  getParentRoute: () => galaxiesRoute,
  path: '/',
  component: () => (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Galaxies</h1>
          <p className="text-muted-foreground">
            Create and manage your galaxies
          </p>
        </div>
      </div>
      <GalaxyTable />
    </div>
  ),
})

const galaxyDetailRoute = createRoute({
  getParentRoute: () => galaxiesRoute,
  path: '$galaxyId',
  component: () => {
    const { galaxyId } = galaxyDetailRoute.useParams()
    return (
      <div>
        <h1 className="text-2xl font-bold">Galaxy Details</h1>
        <p className="text-muted-foreground">Galaxy ID: {galaxyId}</p>
      </div>
    )
  },
})

// Resources routes
const resourcesRoute = createRoute({
  getParentRoute: () => layoutRoute,
  path: 'resources',
  component: () => <Outlet />,
})

const resourcesIndexRoute = createRoute({
  getParentRoute: () => resourcesRoute,
  path: '/',
  component: () => (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Resources</h1>
          <p className="text-muted-foreground">
            Track and verify resources across galaxies
          </p>
        </div>
      </div>
      <ResourceTable />
    </div>
  ),
})

const resourceDetailRoute = createRoute({
  getParentRoute: () => resourcesRoute,
  path: '$resourceId',
  component: () => {
    const { resourceId } = resourceDetailRoute.useParams()
    return (
      <div>
        <h1 className="text-2xl font-bold">Resource Details</h1>
        <p className="text-muted-foreground">Resource ID: {resourceId}</p>
      </div>
    )
  },
})

// Login route (outside layout)
const loginRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: 'login',
  component: () => (
    <div className="flex min-h-screen items-center justify-center">
      <div className="w-full max-w-md space-y-6 p-6">
        <div className="text-center">
          <h1 className="text-2xl font-bold">Sign In</h1>
          <p className="text-muted-foreground">
            Sign in to your Harvester account
          </p>
        </div>
        <p className="text-muted-foreground text-center">
          Login form coming soon...
        </p>
      </div>
    </div>
  ),
})

// 404 route
const notFoundRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: '*',
  component: () => (
    <div className="flex min-h-screen items-center justify-center">
      <div className="text-center">
        <h1 className="text-4xl font-bold">404</h1>
        <p className="text-muted-foreground mt-2">Page not found</p>
      </div>
    </div>
  ),
})

// Build route tree
const routeTree = rootRoute.addChildren([
  layoutRoute.addChildren([
    indexRoute,
    usersRoute.addChildren([usersIndexRoute, userDetailRoute]),
    galaxiesRoute.addChildren([galaxiesIndexRoute, galaxyDetailRoute]),
    resourcesRoute.addChildren([resourcesIndexRoute, resourceDetailRoute]),
  ]),
  loginRoute,
  notFoundRoute,
])

// Create router
export const router = createRouter({
  routeTree,
  defaultPreload: 'intent',
})

// Register router for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

// Simple dashboard card component
function DashboardCard({
  title,
  description,
  href,
  icon,
}: {
  title: string
  description: string
  href: string
  icon: string
}) {
  return (
    <Link
      to={href}
      className="block rounded-lg border bg-card p-6 shadow-sm transition-colors hover:bg-accent"
    >
      <div className="flex items-center gap-4">
        <div className="rounded-md bg-primary/10 p-2">
          <span className="text-2xl">{icon === 'users' ? '👥' : icon === 'globe' ? '🌍' : '📦'}</span>
        </div>
        <div>
          <h3 className="font-semibold">{title}</h3>
          <p className="text-sm text-muted-foreground">{description}</p>
        </div>
      </div>
    </Link>
  )
}
