import { useState } from 'react'
import {
  type ColumnDef,
  type ColumnFiltersState,
  type SortingState,
  type VisibilityState,
  type RowSelectionState,
  type PaginationState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from '@tanstack/react-table'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Skeleton } from '@/components/ui/skeleton'
import { DataTablePagination } from './data-table-pagination'
import { DataTableToolbar } from './data-table-toolbar'

interface DataTableProps<TData, TValue> {
  columns: ColumnDef<TData, TValue>[]
  data: TData[]
  // Server-side pagination
  pageCount?: number
  pagination?: PaginationState
  onPaginationChange?: (pagination: PaginationState) => void
  // Sorting
  sorting?: SortingState
  onSortingChange?: (sorting: SortingState) => void
  // Selection
  enableRowSelection?: boolean
  rowSelection?: RowSelectionState
  onRowSelectionChange?: (selection: RowSelectionState) => void
  // Toolbar
  filterColumn?: string
  filterPlaceholder?: string
  toolbarActions?: React.ReactNode
  // Loading
  isLoading?: boolean
  // Empty state
  emptyMessage?: string
}

export function DataTable<TData, TValue>({
  columns,
  data,
  pageCount,
  pagination,
  onPaginationChange,
  sorting: externalSorting,
  onSortingChange,
  enableRowSelection = false,
  rowSelection: externalRowSelection,
  onRowSelectionChange,
  filterColumn,
  filterPlaceholder,
  toolbarActions,
  isLoading = false,
  emptyMessage = 'No results.',
}: DataTableProps<TData, TValue>) {
  // Internal state for client-side operations
  const [internalSorting, setInternalSorting] = useState<SortingState>([])
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([])
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({})
  const [internalRowSelection, setInternalRowSelection] =
    useState<RowSelectionState>({})
  const [internalPagination, setInternalPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 20,
  })

  // Use external or internal state
  const sorting = externalSorting ?? internalSorting
  const rowSelection = externalRowSelection ?? internalRowSelection
  const currentPagination = pagination ?? internalPagination

  // Determine if pagination/sorting are controlled externally (manual mode)
  const isManualPagination = pageCount !== undefined
  const isManualSorting = externalSorting !== undefined || onSortingChange !== undefined

  const table = useReactTable({
    data,
    columns,
    pageCount: pageCount ?? Math.ceil(data.length / currentPagination.pageSize),
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
      pagination: currentPagination,
    },
    enableRowSelection,
    onRowSelectionChange: (updater) => {
      const newSelection =
        typeof updater === 'function' ? updater(rowSelection) : updater
      if (onRowSelectionChange) {
        onRowSelectionChange(newSelection)
      } else {
        setInternalRowSelection(newSelection)
      }
    },
    onSortingChange: (updater) => {
      const newSorting =
        typeof updater === 'function' ? updater(sorting) : updater
      if (onSortingChange) {
        onSortingChange(newSorting)
      } else {
        setInternalSorting(newSorting)
      }
    },
    onColumnFiltersChange: setColumnFilters,
    onColumnVisibilityChange: setColumnVisibility,
    onPaginationChange: (updater) => {
      const newPagination =
        typeof updater === 'function' ? updater(currentPagination) : updater
      if (onPaginationChange) {
        onPaginationChange(newPagination)
      } else {
        setInternalPagination(newPagination)
      }
    },
    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getPaginationRowModel: isManualPagination ? undefined : getPaginationRowModel(),
    getSortedRowModel: isManualSorting ? undefined : getSortedRowModel(),
    manualPagination: isManualPagination,
    manualSorting: isManualSorting,
  })

  return (
    <div className="space-y-4">
      <DataTableToolbar
        table={table}
        filterColumn={filterColumn}
        filterPlaceholder={filterPlaceholder}
        actions={toolbarActions}
      />
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead key={header.id} colSpan={header.colSpan}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  )
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {isLoading ? (
              // Loading skeleton
              Array.from({ length: currentPagination.pageSize }).map((_, i) => (
                <TableRow key={i}>
                  {columns.map((_, j) => (
                    <TableCell key={j}>
                      <Skeleton className="h-6 w-full" />
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && 'selected'}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  {emptyMessage}
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <DataTablePagination table={table} />
    </div>
  )
}
