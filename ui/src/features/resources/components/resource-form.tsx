import { useEffect } from 'react'
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
import { createResourceSchema, updateResourceSchema, type CreateResourceFormData, type UpdateResourceFormData } from '../schemas'
import type { Resource } from '../types'
import { RESOURCE_STATS } from '../types'
import { useUsers } from '@/features/users'
import { useGalaxies } from '@/features/galaxies'

interface ResourceFormProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  resource?: Resource | null
  onSubmit: (data: CreateResourceFormData | UpdateResourceFormData) => void
  isLoading?: boolean
}

// Create form component
function CreateResourceForm({
  onSubmit,
  isLoading,
  onCancel,
  usersData,
  galaxiesData,
}: {
  onSubmit: (data: CreateResourceFormData) => void
  isLoading?: boolean
  onCancel: () => void
  usersData: { items: Array<{ id: string; name: string; email: string }> } | undefined
  galaxiesData: { items: Array<{ id: string; name: string }> } | undefined
}) {
  const form = useForm<CreateResourceFormData>({
    resolver: zodResolver(createResourceSchema),
    defaultValues: {
      name: '',
      galaxyID: '',
      addedUserID: '',
      resourceType: '',
      cr: 0, cd: 0, dr: 0, fl: 0, hr: 0, ma: 0, pe: 0, oq: 0, sr: 0, ut: 0, er: 0,
    },
  })

  const handleSubmit = form.handleSubmit(onSubmit)

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="space-y-4">
        <h4 className="font-medium text-sm text-muted-foreground">Basic Information</h4>
        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="name">Name</Label>
            <Input id="name" {...form.register('name')} placeholder="Resource name" />
            {form.formState.errors.name && (
              <p className="text-sm text-destructive">{form.formState.errors.name.message}</p>
            )}
          </div>
          <div className="space-y-2">
            <Label htmlFor="resourceType">Type</Label>
            <Input id="resourceType" {...form.register('resourceType')} placeholder="e.g., Copper, Iron..." />
            {form.formState.errors.resourceType && (
              <p className="text-sm text-destructive">{form.formState.errors.resourceType.message}</p>
            )}
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="galaxyID">Galaxy</Label>
            <Select
              value={form.watch('galaxyID')}
              onValueChange={(value) => form.setValue('galaxyID', value, { shouldValidate: true })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Select a galaxy" />
              </SelectTrigger>
              <SelectContent>
                {galaxiesData?.items.map((galaxy) => (
                  <SelectItem key={galaxy.id} value={galaxy.id}>{galaxy.name}</SelectItem>
                ))}
              </SelectContent>
            </Select>
            {form.formState.errors.galaxyID && (
              <p className="text-sm text-destructive">{form.formState.errors.galaxyID.message}</p>
            )}
          </div>
          <div className="space-y-2">
            <Label htmlFor="addedUserID">Added By</Label>
            <Select
              value={form.watch('addedUserID')}
              onValueChange={(value) => form.setValue('addedUserID', value, { shouldValidate: true })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Select a user" />
              </SelectTrigger>
              <SelectContent>
                {usersData?.items.map((user) => (
                  <SelectItem key={user.id} value={user.id}>{user.name} ({user.email})</SelectItem>
                ))}
              </SelectContent>
            </Select>
            {form.formState.errors.addedUserID && (
              <p className="text-sm text-destructive">{form.formState.errors.addedUserID.message}</p>
            )}
          </div>
        </div>
      </div>

      <div className="space-y-4">
        <h4 className="font-medium text-sm text-muted-foreground">Stats (0-1000)</h4>
        <div className="grid grid-cols-3 sm:grid-cols-4 gap-3">
          {RESOURCE_STATS.map((stat) => (
            <div key={stat.key} className="space-y-1">
              <Label htmlFor={stat.key} className="text-xs">
                {stat.abbr}
                <span className="sr-only"> - {stat.label}</span>
              </Label>
              <Input
                id={stat.key}
                type="number"
                min={0}
                max={1000}
                {...form.register(stat.key, { valueAsNumber: true })}
                className="h-8 text-sm"
              />
            </div>
          ))}
        </div>
      </div>

      <DialogFooter>
        <Button type="button" variant="outline" onClick={onCancel}>Cancel</Button>
        <Button type="submit" disabled={isLoading}>
          {isLoading ? 'Saving...' : 'Create Resource'}
        </Button>
      </DialogFooter>
    </form>
  )
}

// Edit form component
function EditResourceForm({
  resource,
  onSubmit,
  isLoading,
  onCancel,
  galaxiesData,
}: {
  resource: Resource
  onSubmit: (data: UpdateResourceFormData) => void
  isLoading?: boolean
  onCancel: () => void
  galaxiesData: { items: Array<{ id: string; name: string }> } | undefined
}) {
  const form = useForm<UpdateResourceFormData>({
    resolver: zodResolver(updateResourceSchema),
    defaultValues: {
      name: resource.name,
      galaxyID: resource.galaxyID,
      resourceType: resource.resourceType,
      verified: resource.verified,
      cr: resource.cr, cd: resource.cd, dr: resource.dr, fl: resource.fl, hr: resource.hr,
      ma: resource.ma, pe: resource.pe, oq: resource.oq, sr: resource.sr, ut: resource.ut, er: resource.er,
    },
  })

  // Reset when resource changes
  useEffect(() => {
    form.reset({
      name: resource.name,
      galaxyID: resource.galaxyID,
      resourceType: resource.resourceType,
      verified: resource.verified,
      cr: resource.cr, cd: resource.cd, dr: resource.dr, fl: resource.fl, hr: resource.hr,
      ma: resource.ma, pe: resource.pe, oq: resource.oq, sr: resource.sr, ut: resource.ut, er: resource.er,
    })
  }, [resource, form])

  const handleSubmit = form.handleSubmit(onSubmit)

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      <div className="space-y-4">
        <h4 className="font-medium text-sm text-muted-foreground">Basic Information</h4>
        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="name">Name</Label>
            <Input id="name" {...form.register('name')} placeholder="Resource name" />
            {form.formState.errors.name && (
              <p className="text-sm text-destructive">{form.formState.errors.name.message}</p>
            )}
          </div>
          <div className="space-y-2">
            <Label htmlFor="resourceType">Type</Label>
            <Input id="resourceType" {...form.register('resourceType')} placeholder="e.g., Copper, Iron..." />
            {form.formState.errors.resourceType && (
              <p className="text-sm text-destructive">{form.formState.errors.resourceType.message}</p>
            )}
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div className="space-y-2">
            <Label htmlFor="galaxyID">Galaxy</Label>
            <Select
              value={form.watch('galaxyID')}
              onValueChange={(value) => form.setValue('galaxyID', value, { shouldValidate: true })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Select a galaxy" />
              </SelectTrigger>
              <SelectContent>
                {galaxiesData?.items.map((galaxy) => (
                  <SelectItem key={galaxy.id} value={galaxy.id}>{galaxy.name}</SelectItem>
                ))}
              </SelectContent>
            </Select>
            {form.formState.errors.galaxyID && (
              <p className="text-sm text-destructive">{form.formState.errors.galaxyID.message}</p>
            )}
          </div>
          <div className="flex items-center justify-between pt-6">
            <Label htmlFor="verified">Verified</Label>
            <Switch
              id="verified"
              checked={form.watch('verified') ?? false}
              onCheckedChange={(checked) => form.setValue('verified', checked)}
            />
          </div>
        </div>
      </div>

      <div className="space-y-4">
        <h4 className="font-medium text-sm text-muted-foreground">Stats (0-1000)</h4>
        <div className="grid grid-cols-3 sm:grid-cols-4 gap-3">
          {RESOURCE_STATS.map((stat) => (
            <div key={stat.key} className="space-y-1">
              <Label htmlFor={stat.key} className="text-xs">
                {stat.abbr}
                <span className="sr-only"> - {stat.label}</span>
              </Label>
              <Input
                id={stat.key}
                type="number"
                min={0}
                max={1000}
                {...form.register(stat.key, { valueAsNumber: true })}
                className="h-8 text-sm"
              />
            </div>
          ))}
        </div>
      </div>

      <DialogFooter>
        <Button type="button" variant="outline" onClick={onCancel}>Cancel</Button>
        <Button type="submit" disabled={isLoading}>
          {isLoading ? 'Saving...' : 'Save Changes'}
        </Button>
      </DialogFooter>
    </form>
  )
}

export function ResourceForm({ open, onOpenChange, resource, onSubmit, isLoading }: ResourceFormProps) {
  const isEditing = !!resource

  // Fetch users and galaxies for selects
  const { data: usersData } = useUsers({ rows: 100 })
  const { data: galaxiesData } = useGalaxies({ rows: 100 })

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>{isEditing ? 'Edit Resource' : 'Create Resource'}</DialogTitle>
          <DialogDescription>
            {isEditing
              ? 'Make changes to the resource.'
              : 'Add a new resource to the system.'}
          </DialogDescription>
        </DialogHeader>

        {isEditing && resource ? (
          <EditResourceForm
            resource={resource}
            onSubmit={onSubmit}
            isLoading={isLoading}
            onCancel={() => onOpenChange(false)}
            galaxiesData={galaxiesData}
          />
        ) : (
          <CreateResourceForm
            onSubmit={onSubmit}
            isLoading={isLoading}
            onCancel={() => onOpenChange(false)}
            usersData={usersData}
            galaxiesData={galaxiesData}
          />
        )}
      </DialogContent>
    </Dialog>
  )
}
