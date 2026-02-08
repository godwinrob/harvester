import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Switch } from '@/components/ui/switch'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { createGalaxySchema, updateGalaxySchema } from '../schemas'
import type { CreateGalaxyFormData, UpdateGalaxyFormData } from '../schemas'
import type { Galaxy } from '../types'
import { useUsers } from '@/features/users'

interface GalaxyFormProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  galaxy?: Galaxy | null
  onSubmit: (data: CreateGalaxyFormData | UpdateGalaxyFormData) => void
  isLoading?: boolean
}

export function GalaxyForm({ open, onOpenChange, galaxy, onSubmit, isLoading }: GalaxyFormProps) {
  const isEditing = !!galaxy

  // Fetch users for owner select
  const { data: usersData } = useUsers({ rows: 100 })

  const form = useForm<CreateGalaxyFormData | UpdateGalaxyFormData>({
    resolver: zodResolver(isEditing ? updateGalaxySchema : createGalaxySchema),
    defaultValues: isEditing
      ? {
          name: galaxy.name,
          ownerUserID: galaxy.ownerUserID,
          enabled: galaxy.enabled,
        }
      : {
          name: '',
          ownerUserID: '',
        },
  })

  const handleSubmit = form.handleSubmit((data) => {
    onSubmit(data)
  })

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>{isEditing ? 'Edit Galaxy' : 'Create Galaxy'}</DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Make changes to the galaxy.'
              : 'Add a new galaxy to the system.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="name">Name</Label>
            <Input
              id="name"
              {...form.register('name')}
              placeholder="Galaxy name"
            />
            {form.formState.errors.name && (
              <p className="text-sm text-destructive">
                {form.formState.errors.name.message}
              </p>
            )}
          </div>

          <div className="space-y-2">
            <Label htmlFor="ownerUserID">Owner</Label>
            <Select
              value={form.watch('ownerUserID')}
              onValueChange={(value) => form.setValue('ownerUserID', value, { shouldValidate: true })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Select an owner" />
              </SelectTrigger>
              <SelectContent>
                {usersData?.items.map((user) => (
                  <SelectItem key={user.id} value={user.id}>
                    {user.name} ({user.email})
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {form.formState.errors.ownerUserID && (
              <p className="text-sm text-destructive">
                {form.formState.errors.ownerUserID.message}
              </p>
            )}
          </div>

          {isEditing && (
            <div className="flex items-center justify-between">
              <Label htmlFor="enabled">Galaxy Enabled</Label>
              <Switch
                id="enabled"
                checked={form.watch('enabled' as keyof UpdateGalaxyFormData) as boolean}
                onCheckedChange={(checked) =>
                  form.setValue('enabled' as keyof UpdateGalaxyFormData, checked)
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
              {isLoading ? 'Saving...' : isEditing ? 'Save Changes' : 'Create Galaxy'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
