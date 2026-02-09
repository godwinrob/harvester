import { useState, useMemo, useEffect } from 'react'
import type { PaginationState, SortingState, RowSelectionState } from '@tanstack/react-table'
import { Plus, Trash2, X } from 'lucide-react'
import { toast } from 'sonner'
import { DataTable } from '@/components/ui/data-table'
import { Button } from '@/components/ui/button'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { useResources, useCreateResource, useUpdateResource, useDeleteResource, useBulkDeleteResources } from '../hooks'
import { getResourceColumns } from './resource-columns'
import { ResourceForm } from './resource-form'
import { DeleteResourceDialog, BulkDeleteResourcesDialog } from './delete-resource-dialog'
import type { Resource, ResourceListParams } from '../types'
import type { CreateResourceFormData, UpdateResourceFormData } from '../schemas'
import { ApiError } from '@/lib/api'
import { resourceTypesApi } from '@/features/resource-types/api/resource-types.api'
import { resourceGroupsApi } from '@/features/resource-types/api/resource-groups.api'
import type { ResourceType, ResourceGroup } from '@/features/resource-types/types'

export function ResourceTable() {
  // Resource groups fetched from API
  const [resourceGroups, setResourceGroups] = useState<ResourceGroup[]>([])

  // Resource types fetched from API
  const [allResourceTypes, setAllResourceTypes] = useState<ResourceType[]>([])

  useEffect(() => {
    resourceGroupsApi.list({ rows: 2000 }).then((response) => {
      setResourceGroups(response.items)
    })
    resourceTypesApi.list({ rows: 2000, enterable: true }).then((response) => {
      setAllResourceTypes(response.items)
    })
  }, [])

  // Pagination state
  const [pagination, setPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 20,
  })

  // Sorting state
  const [sorting, setSorting] = useState<SortingState>([])

  // Row selection state
  const [rowSelection, setRowSelection] = useState<RowSelectionState>({})

  // Filter state
  const [resourceGroupFilter, setResourceGroupFilter] = useState<string>('')
  const [resourceTypeFilter, setResourceTypeFilter] = useState<string>('')

  // Filter resource types by selected group
  const filteredResourceTypes = useMemo(() => {
    const types = resourceGroupFilter
      ? allResourceTypes.filter((rt) => rt.resourceGroup === resourceGroupFilter)
      : allResourceTypes
    // Group by category for the dropdown
    const grouped: Record<string, ResourceType[]> = {}
    for (const rt of types) {
      const category = rt.resourceCategory || 'Other'
      if (!grouped[category]) grouped[category] = []
      grouped[category].push(rt)
    }
    return grouped
  }, [allResourceTypes, resourceGroupFilter])

  // Clear type filter if it no longer matches the selected group
  useEffect(() => {
    if (resourceGroupFilter && resourceTypeFilter) {
      const typeStillValid = allResourceTypes.some(
        (rt) => rt.resourceType === resourceTypeFilter && rt.resourceGroup === resourceGroupFilter
      )
      if (!typeStillValid) {
        setResourceTypeFilter('')
      }
    }
  }, [resourceGroupFilter, resourceTypeFilter, allResourceTypes])

  // Dialog states
  const [formOpen, setFormOpen] = useState(false)
  const [editingResource, setEditingResource] = useState<Resource | null>(null)
  const [deletingResource, setDeletingResource] = useState<Resource | null>(null)
  const [bulkDeleteOpen, setBulkDeleteOpen] = useState(false)

  // Build query params
  const queryParams: ResourceListParams = useMemo(
    () => ({
      page: pagination.pageIndex + 1,
      rows: pagination.pageSize,
      orderBy: sorting[0]
        ? {
            field: sorting[0].id,
            direction: sorting[0].desc ? 'DESC' : 'ASC',
          }
        : undefined,
      resourceGroup: resourceGroupFilter || undefined,
      resourceType: resourceTypeFilter || undefined,
    }),
    [pagination, sorting, resourceGroupFilter, resourceTypeFilter]
  )

  // Queries and mutations
  const { data, isLoading } = useResources(queryParams)
  const createResource = useCreateResource()
  const updateResource = useUpdateResource()
  const deleteResource = useDeleteResource()
  const bulkDeleteResources = useBulkDeleteResources()

  // Get selected resource IDs
  const selectedResourceIds = useMemo(() => {
    if (!data?.items) return []
    return Object.keys(rowSelection)
      .filter((key) => rowSelection[key])
      .map((index) => data.items[parseInt(index)]?.id)
      .filter(Boolean)
  }, [rowSelection, data?.items])

  // Column definitions with handlers
  const columns = useMemo(
    () =>
      getResourceColumns({
        onEdit: (resource) => {
          setEditingResource(resource)
          setFormOpen(true)
        },
        onDelete: (resource) => {
          setDeletingResource(resource)
        },
      }),
    []
  )

  // Handle form submit
  const handleFormSubmit = async (formData: CreateResourceFormData | UpdateResourceFormData) => {
    try {
      if (editingResource) {
        await updateResource.mutateAsync({ id: editingResource.id, data: formData })
        toast.success('Resource updated successfully')
      } else {
        await createResource.mutateAsync(formData as CreateResourceFormData)
        toast.success('Resource created successfully')
      }
      setFormOpen(false)
      setEditingResource(null)
    } catch (error) {
      if (error instanceof ApiError) {
        toast.error(error.message)
      } else {
        toast.error('An error occurred')
      }
    }
  }

  // Handle single delete
  const handleDelete = async () => {
    if (!deletingResource) return
    try {
      await deleteResource.mutateAsync(deletingResource.id)
      toast.success('Resource deleted successfully')
      setDeletingResource(null)
    } catch (error) {
      if (error instanceof ApiError) {
        toast.error(error.message)
      } else {
        toast.error('An error occurred')
      }
    }
  }

  // Handle bulk delete
  const handleBulkDelete = async () => {
    try {
      await bulkDeleteResources.mutateAsync(selectedResourceIds)
      toast.success(`${selectedResourceIds.length} resources deleted successfully`)
      setRowSelection({})
      setBulkDeleteOpen(false)
    } catch (error) {
      if (error instanceof ApiError) {
        toast.error(error.message)
      } else {
        toast.error('An error occurred')
      }
    }
  }

  // Calculate page count
  const pageCount = data ? Math.ceil(data.total / pagination.pageSize) : 0

  return (
    <>
      <DataTable
        columns={columns}
        data={data?.items ?? []}
        pageCount={pageCount}
        pagination={pagination}
        onPaginationChange={setPagination}
        sorting={sorting}
        onSortingChange={setSorting}
        enableRowSelection
        rowSelection={rowSelection}
        onRowSelectionChange={setRowSelection}
        filterColumn="name"
        filterPlaceholder="Filter by name..."
        isLoading={isLoading}
        emptyMessage="No resources found."
        toolbarActions={
          <div className="flex gap-2 items-center">
            {/* Resource Group Filter */}
            <div className="flex items-center gap-1">
              <Select
                value={resourceGroupFilter}
                onValueChange={setResourceGroupFilter}
              >
                <SelectTrigger className="w-[200px] h-8">
                  <SelectValue placeholder="Filter by group..." />
                </SelectTrigger>
                <SelectContent className="max-h-[400px]">
                  {resourceGroups.map((rg) => (
                    <SelectItem key={rg.resourceGroup} value={rg.resourceGroup}>
                      {rg.groupName}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {resourceGroupFilter && (
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => setResourceGroupFilter('')}
                  className="h-8 px-2"
                  title="Clear group filter"
                >
                  <X className="h-4 w-4" />
                </Button>
              )}
            </div>

            {/* Resource Type Filter */}
            <div className="flex items-center gap-1">
              <Select
                value={resourceTypeFilter}
                onValueChange={setResourceTypeFilter}
              >
                <SelectTrigger className="w-[240px] h-8">
                  <SelectValue placeholder="Filter by type..." />
                </SelectTrigger>
                <SelectContent className="max-h-[400px]">
                  {Object.entries(filteredResourceTypes).map(([category, types]) => (
                    <SelectGroup key={category}>
                      <SelectLabel>{category}</SelectLabel>
                      {types.map((rt) => (
                        <SelectItem key={rt.resourceType} value={rt.resourceType}>
                          {rt.resourceTypeName}
                        </SelectItem>
                      ))}
                    </SelectGroup>
                  ))}
                </SelectContent>
              </Select>
              {resourceTypeFilter && (
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => setResourceTypeFilter('')}
                  className="h-8 px-2"
                  title="Clear type filter"
                >
                  <X className="h-4 w-4" />
                </Button>
              )}
            </div>

            {selectedResourceIds.length > 0 && (
              <Button
                variant="destructive"
                size="sm"
                onClick={() => setBulkDeleteOpen(true)}
              >
                <Trash2 className="mr-2 h-4 w-4" />
                Delete ({selectedResourceIds.length})
              </Button>
            )}
            <Button
              size="sm"
              onClick={() => {
                setEditingResource(null)
                setFormOpen(true)
              }}
            >
              <Plus className="mr-2 h-4 w-4" />
              Add Resource
            </Button>
          </div>
        }
      />

      {/* Create/Edit Form Dialog */}
      <ResourceForm
        open={formOpen}
        onOpenChange={(open) => {
          setFormOpen(open)
          if (!open) setEditingResource(null)
        }}
        resource={editingResource}
        onSubmit={handleFormSubmit}
        isLoading={createResource.isPending || updateResource.isPending}
      />

      {/* Delete Confirmation Dialog */}
      <DeleteResourceDialog
        open={!!deletingResource}
        onOpenChange={(open) => !open && setDeletingResource(null)}
        resource={deletingResource}
        onConfirm={handleDelete}
        isLoading={deleteResource.isPending}
      />

      {/* Bulk Delete Confirmation Dialog */}
      <BulkDeleteResourcesDialog
        open={bulkDeleteOpen}
        onOpenChange={setBulkDeleteOpen}
        count={selectedResourceIds.length}
        onConfirm={handleBulkDelete}
        isLoading={bulkDeleteResources.isPending}
      />
    </>
  )
}
