import type { ColumnDef } from '@tanstack/react-table'
import { MoreHorizontal, Pencil, Trash2, CheckCircle2, XCircle } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Badge } from '@/components/ui/badge'
import { DataTableColumnHeader } from '@/components/ui/data-table'
import type { Resource } from '../types'
import { RESOURCE_STATS } from '../types'

interface ColumnOptions {
  onEdit: (resource: Resource) => void
  onDelete: (resource: Resource) => void
}

export function getResourceColumns({ onEdit, onDelete }: ColumnOptions): ColumnDef<Resource>[] {
  return [
    {
      id: 'select',
      header: ({ table }) => (
        <Checkbox
          checked={
            table.getIsAllPageRowsSelected() ||
            (table.getIsSomePageRowsSelected() && 'indeterminate')
          }
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
        />
      ),
      cell: ({ row }) => (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
        />
      ),
      enableSorting: false,
      enableHiding: false,
    },
    {
      accessorKey: 'name',
      header: ({ column }) => <DataTableColumnHeader column={column} title="Name" />,
      cell: ({ row }) => (
        <div className="font-medium">{row.getValue('name')}</div>
      ),
    },
    {
      accessorKey: 'resourceType',
      header: ({ column }) => <DataTableColumnHeader column={column} title="Type" />,
      cell: ({ row }) => (
        <Badge variant="outline">{row.getValue('resourceType')}</Badge>
      ),
    },
    {
      accessorKey: 'oq',
      header: ({ column }) => <DataTableColumnHeader column={column} title="OQ" />,
      cell: ({ row }) => {
        const value = row.original.oq
        return (
          <div className="text-center">
            <span className={value >= 900 ? 'text-green-600 font-semibold' : value >= 700 ? 'text-yellow-600' : ''}>
              {value}
            </span>
          </div>
        )
      },
    },
    {
      id: 'stats',
      header: 'Stats',
      cell: ({ row }) => {
        const resource = row.original
        // Show abbreviated stats in a compact format
        return (
          <div className="flex gap-1 flex-wrap max-w-[300px]">
            {RESOURCE_STATS.filter(stat => stat.key !== 'oq').slice(0, 5).map((stat) => (
              <Badge key={stat.key} variant="secondary" className="text-xs">
                {stat.abbr}: {resource[stat.key]}
              </Badge>
            ))}
          </div>
        )
      },
      enableSorting: false,
    },
    {
      accessorKey: 'verified',
      header: ({ column }) => <DataTableColumnHeader column={column} title="Verified" />,
      cell: ({ row }) => {
        const verified = row.getValue('verified') as boolean
        return verified ? (
          <CheckCircle2 className="h-4 w-4 text-green-600" />
        ) : (
          <XCircle className="h-4 w-4 text-muted-foreground" />
        )
      },
    },
    {
      accessorKey: 'unavailableAt',
      header: ({ column }) => <DataTableColumnHeader column={column} title="Status" />,
      cell: ({ row }) => {
        const unavailableAt = row.getValue('unavailableAt') as string | null
        return unavailableAt ? (
          <Badge variant="destructive">Unavailable</Badge>
        ) : (
          <Badge variant="default">Available</Badge>
        )
      },
    },
    {
      accessorKey: 'addedAtDate',
      header: ({ column }) => <DataTableColumnHeader column={column} title="Added" />,
      cell: ({ row }) => {
        const date = new Date(row.getValue('addedAtDate') as string)
        return <div className="text-muted-foreground">{date.toLocaleDateString()}</div>
      },
    },
    {
      id: 'actions',
      cell: ({ row }) => {
        const resource = row.original

        return (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" className="h-8 w-8 p-0">
                <span className="sr-only">Open menu</span>
                <MoreHorizontal className="h-4 w-4" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Actions</DropdownMenuLabel>
              <DropdownMenuItem onClick={() => navigator.clipboard.writeText(resource.id)}>
                Copy ID
              </DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem onClick={() => onEdit(resource)}>
                <Pencil className="mr-2 h-4 w-4" />
                Edit
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={() => onDelete(resource)}
                className="text-destructive"
              >
                <Trash2 className="mr-2 h-4 w-4" />
                Delete
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        )
      },
    },
  ]
}
