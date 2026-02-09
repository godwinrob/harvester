import { useState, useMemo } from 'react'
import type { PaginationState, SortingState, RowSelectionState } from '@tanstack/react-table'
import { Plus, Trash2 } from 'lucide-react'
import { toast } from 'sonner'
import { DataTable } from '@/components/ui/data-table'
import { Button } from '@/components/ui/button'
import { useGalaxies, useCreateGalaxy, useUpdateGalaxy, useDeleteGalaxy, useBulkDeleteGalaxies } from '../hooks'
import { getGalaxyColumns } from './galaxy-columns'
import { GalaxyForm } from './galaxy-form'
import { DeleteGalaxyDialog, BulkDeleteGalaxiesDialog } from './delete-galaxy-dialog'
import type { Galaxy, GalaxyListParams } from '../types'
import type { CreateGalaxyFormData, UpdateGalaxyFormData } from '../schemas'
import { ApiError } from '@/lib/api'

export function GalaxyTable() {
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
  const [editingGalaxy, setEditingGalaxy] = useState<Galaxy | null>(null)
  const [deletingGalaxy, setDeletingGalaxy] = useState<Galaxy | null>(null)
  const [bulkDeleteOpen, setBulkDeleteOpen] = useState(false)

  // Build query params
  const queryParams: GalaxyListParams = useMemo(
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
  const { data, isLoading } = useGalaxies(queryParams)
  const createGalaxy = useCreateGalaxy()
  const updateGalaxy = useUpdateGalaxy()
  const deleteGalaxy = useDeleteGalaxy()
  const bulkDeleteGalaxies = useBulkDeleteGalaxies()

  // Get selected galaxy IDs
  const selectedGalaxyIds = useMemo(() => {
    if (!data?.items) return []
    return Object.keys(rowSelection)
      .filter((key) => rowSelection[key])
      .map((index) => data.items[parseInt(index)]?.id)
      .filter(Boolean)
  }, [rowSelection, data?.items])

  // Column definitions with handlers
  const columns = useMemo(
    () =>
      getGalaxyColumns({
        onEdit: (galaxy) => {
          setEditingGalaxy(galaxy)
          setFormOpen(true)
        },
        onDelete: (galaxy) => {
          setDeletingGalaxy(galaxy)
        },
      }),
    []
  )

  // Handle form submit
  const handleFormSubmit = async (formData: CreateGalaxyFormData | UpdateGalaxyFormData) => {
    try {
      if (editingGalaxy) {
        await updateGalaxy.mutateAsync({ id: editingGalaxy.id, data: formData })
        toast.success('Galaxy updated successfully')
      } else {
        await createGalaxy.mutateAsync(formData as CreateGalaxyFormData)
        toast.success('Galaxy created successfully')
      }
      setFormOpen(false)
      setEditingGalaxy(null)
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
    if (!deletingGalaxy) return
    try {
      await deleteGalaxy.mutateAsync(deletingGalaxy.id)
      toast.success('Galaxy deleted successfully')
      setDeletingGalaxy(null)
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
      await bulkDeleteGalaxies.mutateAsync(selectedGalaxyIds)
      toast.success(`${selectedGalaxyIds.length} galaxies deleted successfully`)
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
        emptyMessage="No galaxies found."
        toolbarActions={
          <div className="flex gap-2">
            {selectedGalaxyIds.length > 0 && (
              <Button
                variant="destructive"
                size="sm"
                onClick={() => setBulkDeleteOpen(true)}
              >
                <Trash2 className="mr-2 h-4 w-4" />
                Delete ({selectedGalaxyIds.length})
              </Button>
            )}
            <Button
              size="sm"
              onClick={() => {
                setEditingGalaxy(null)
                setFormOpen(true)
              }}
            >
              <Plus className="mr-2 h-4 w-4" />
              Add Galaxy
            </Button>
          </div>
        }
      />

      {/* Create/Edit Form Dialog */}
      <GalaxyForm
        open={formOpen}
        onOpenChange={(open) => {
          setFormOpen(open)
          if (!open) setEditingGalaxy(null)
        }}
        galaxy={editingGalaxy}
        onSubmit={handleFormSubmit}
        isLoading={createGalaxy.isPending || updateGalaxy.isPending}
      />

      {/* Delete Confirmation Dialog */}
      <DeleteGalaxyDialog
        open={!!deletingGalaxy}
        onOpenChange={(open) => !open && setDeletingGalaxy(null)}
        galaxy={deletingGalaxy}
        onConfirm={handleDelete}
        isLoading={deleteGalaxy.isPending}
      />

      {/* Bulk Delete Confirmation Dialog */}
      <BulkDeleteGalaxiesDialog
        open={bulkDeleteOpen}
        onOpenChange={setBulkDeleteOpen}
        count={selectedGalaxyIds.length}
        onConfirm={handleBulkDelete}
        isLoading={bulkDeleteGalaxies.isPending}
      />
    </>
  )
}
