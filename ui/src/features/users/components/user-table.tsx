import { useState, useMemo } from 'react'
import type { PaginationState, SortingState, RowSelectionState } from '@tanstack/react-table'
import { Plus, Trash2 } from 'lucide-react'
import { toast } from 'sonner'
import { DataTable } from '@/components/ui/data-table'
import { Button } from '@/components/ui/button'
import { useUsers, useCreateUser, useUpdateUser, useDeleteUser, useBulkDeleteUsers } from '../hooks'
import { getUserColumns } from './user-columns'
import { UserForm } from './user-form'
import { DeleteUserDialog, BulkDeleteUsersDialog } from './delete-user-dialog'
import type { User, UserListParams } from '../types'
import type { CreateUserFormData, UpdateUserFormData } from '../schemas'
import { ApiError } from '@/lib/api'

export function UserTable() {
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
  const [editingUser, setEditingUser] = useState<User | null>(null)
  const [deletingUser, setDeletingUser] = useState<User | null>(null)
  const [bulkDeleteOpen, setBulkDeleteOpen] = useState(false)

  // Build query params
  const queryParams: UserListParams = useMemo(
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
  const { data, isLoading } = useUsers(queryParams)
  const createUser = useCreateUser()
  const updateUser = useUpdateUser()
  const deleteUser = useDeleteUser()
  const bulkDeleteUsers = useBulkDeleteUsers()

  // Get selected user IDs
  const selectedUserIds = useMemo(() => {
    if (!data?.items) return []
    return Object.keys(rowSelection)
      .filter((key) => rowSelection[key])
      .map((index) => data.items[parseInt(index)]?.id)
      .filter(Boolean)
  }, [rowSelection, data?.items])

  // Column definitions with handlers
  const columns = useMemo(
    () =>
      getUserColumns({
        onEdit: (user) => {
          setEditingUser(user)
          setFormOpen(true)
        },
        onDelete: (user) => {
          setDeletingUser(user)
        },
      }),
    []
  )

  // Handle form submit
  const handleFormSubmit = async (formData: CreateUserFormData | UpdateUserFormData) => {
    try {
      if (editingUser) {
        await updateUser.mutateAsync({ id: editingUser.id, data: formData })
        toast.success('User updated successfully')
      } else {
        await createUser.mutateAsync(formData as CreateUserFormData)
        toast.success('User created successfully')
      }
      setFormOpen(false)
      setEditingUser(null)
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
    if (!deletingUser) return
    try {
      await deleteUser.mutateAsync(deletingUser.id)
      toast.success('User deleted successfully')
      setDeletingUser(null)
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
      await bulkDeleteUsers.mutateAsync(selectedUserIds)
      toast.success(`${selectedUserIds.length} users deleted successfully`)
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
        emptyMessage="No users found."
        toolbarActions={
          <div className="flex gap-2">
            {selectedUserIds.length > 0 && (
              <Button
                variant="destructive"
                size="sm"
                onClick={() => setBulkDeleteOpen(true)}
              >
                <Trash2 className="mr-2 h-4 w-4" />
                Delete ({selectedUserIds.length})
              </Button>
            )}
            <Button
              size="sm"
              onClick={() => {
                setEditingUser(null)
                setFormOpen(true)
              }}
            >
              <Plus className="mr-2 h-4 w-4" />
              Add User
            </Button>
          </div>
        }
      />

      {/* Create/Edit Form Dialog */}
      <UserForm
        open={formOpen}
        onOpenChange={(open) => {
          setFormOpen(open)
          if (!open) setEditingUser(null)
        }}
        user={editingUser}
        onSubmit={handleFormSubmit}
        isLoading={createUser.isPending || updateUser.isPending}
      />

      {/* Delete Confirmation Dialog */}
      <DeleteUserDialog
        open={!!deletingUser}
        onOpenChange={(open) => !open && setDeletingUser(null)}
        user={deletingUser}
        onConfirm={handleDelete}
        isLoading={deleteUser.isPending}
      />

      {/* Bulk Delete Confirmation Dialog */}
      <BulkDeleteUsersDialog
        open={bulkDeleteOpen}
        onOpenChange={setBulkDeleteOpen}
        count={selectedUserIds.length}
        onConfirm={handleBulkDelete}
        isLoading={bulkDeleteUsers.isPending}
      />
    </>
  )
}
