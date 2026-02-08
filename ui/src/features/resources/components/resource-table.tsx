import { useState, useMemo } from 'react'
import type { PaginationState, SortingState, RowSelectionState } from '@tanstack/react-table'
import { Plus, Trash2 } from 'lucide-react'
import { toast } from 'sonner'
import { DataTable } from '@/components/ui/data-table'
import { Button } from '@/components/ui/button'
import { useResources, useCreateResource, useUpdateResource, useDeleteResource, useBulkDeleteResources } from '../hooks'
import { getResourceColumns } from './resource-columns'
import { ResourceForm } from './resource-form'
import { DeleteResourceDialog, BulkDeleteResourcesDialog } from './delete-resource-dialog'
import type { Resource, ResourceListParams } from '../types'
import type { CreateResourceFormData, UpdateResourceFormData } from '../schemas'
import { ApiError } from '@/lib/api'

export function ResourceTable() {
  // Pagination state
  const [pagination, setPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 20,
  })

  // Sorting state
  const [sorting, setSorting] = useState<SortingState>([])

  // Row selection state
  const [rowSelection, setRowSelection] = useState<RowSelectionState>({})

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
    }),
    [pagination, sorting]
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
          <div className="flex gap-2">
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
