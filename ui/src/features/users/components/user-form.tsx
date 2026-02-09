import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Checkbox } from '@/components/ui/checkbox'
import { Switch } from '@/components/ui/switch'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { createUserSchema, updateUserSchema } from '../schemas'
import type { CreateUserFormData, UpdateUserFormData } from '../schemas'
import type { User, UserRole } from '../types'

interface UserFormProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  user?: User | null
  onSubmit: (data: CreateUserFormData | UpdateUserFormData) => void
  isLoading?: boolean
}

const ROLES: UserRole[] = ['ADMIN', 'USER']

export function UserForm({ open, onOpenChange, user, onSubmit, isLoading }: UserFormProps) {
  const isEditing = !!user

  const form = useForm<CreateUserFormData | UpdateUserFormData>({
    resolver: zodResolver(isEditing ? updateUserSchema : createUserSchema),
    defaultValues: isEditing
      ? {
          name: user.name,
          email: user.email,
          roles: user.roles,
          guild: user.guild ?? '',
          enabled: user.enabled,
        }
      : {
          name: '',
          email: '',
          password: '',
          passwordConfirm: '',
          roles: ['USER'],
          guild: '',
        },
  })

  const handleSubmit = form.handleSubmit((data) => {
    onSubmit(data)
  })

  const selectedRoles = form.watch('roles') ?? []

  const toggleRole = (role: UserRole) => {
    const current = form.getValues('roles') ?? []
    if (current.includes(role)) {
      if (current.length > 1) {
        form.setValue(
          'roles',
          current.filter((r) => r !== role),
          { shouldValidate: true }
        )
      }
    } else {
      form.setValue('roles', [...current, role], { shouldValidate: true })
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>{isEditing ? 'Edit User' : 'Create User'}</DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Make changes to the user account.'
              : 'Add a new user to the system.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Name</Label>
            <Input
              id="name"
              {...form.register('name')}
              placeholder="John Doe"
            />
            {form.formState.errors.name && (
              <p className="text-sm text-destructive">
                {form.formState.errors.name.message}
              </p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              type="email"
              {...form.register('email')}
              placeholder="john@example.com"
            />
            {form.formState.errors.email && (
              <p className="text-sm text-destructive">
                {form.formState.errors.email.message}
              </p>
            )}
          </div>

          {!isEditing && (
            <>
              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  type="password"
                  {...form.register('password' as keyof CreateUserFormData)}
                  placeholder="••••••••"
                />
                {(form.formState.errors as { password?: { message?: string } }).password && (
                  <p className="text-sm text-destructive">
                    {(form.formState.errors as { password?: { message?: string } }).password?.message}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="passwordConfirm">Confirm Password</Label>
                <Input
                  id="passwordConfirm"
                  type="password"
                  {...form.register('passwordConfirm' as keyof CreateUserFormData)}
                  placeholder="••••••••"
                />
                {(form.formState.errors as { passwordConfirm?: { message?: string } }).passwordConfirm && (
                  <p className="text-sm text-destructive">
                    {(form.formState.errors as { passwordConfirm?: { message?: string } }).passwordConfirm?.message}
                  </p>
                )}
              </div>
            </>
          )}

          <div className="space-y-2">
            <Label>Roles</Label>
            <div className="flex gap-4">
              {ROLES.map((role) => (
                <label
                  key={role}
                  className="flex items-center gap-2 cursor-pointer"
                >
                  <Checkbox
                    checked={selectedRoles.includes(role)}
                    onCheckedChange={() => toggleRole(role)}
                  />
                  <span className="text-sm">{role}</span>
                </label>
              ))}
            </div>
            {form.formState.errors.roles && (
              <p className="text-sm text-destructive">
                {form.formState.errors.roles.message}
              </p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="guild">Guild (optional)</Label>
            <Input
              id="guild"
              {...form.register('guild')}
              placeholder="Guild name"
            />
          </div>

          {isEditing && (
            <div className="flex items-center justify-between">
              <Label htmlFor="enabled">Account Enabled</Label>
              <Switch
                id="enabled"
                checked={form.watch('enabled' as keyof UpdateUserFormData) as boolean}
                onCheckedChange={(checked) =>
                  form.setValue('enabled' as keyof UpdateUserFormData, checked)
                }
              />
            </div>
          )}

          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => onOpenChange(false)}
            >
              Cancel
            </Button>
            <Button type="submit" disabled={isLoading}>
              {isLoading ? 'Saving...' : isEditing ? 'Save Changes' : 'Create User'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
