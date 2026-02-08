import type { ColumnDef } from '@tanstack/react-table'
import { MoreHorizontal, Pencil, Trash2 } from 'lucide-react'
import { Badge } from '@/components/ui/badge'
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
import { DataTableColumnHeader } from '@/components/ui/data-table'
import type { Galaxy } from '../types'

interface GalaxyColumnsProps {
  onEdit: (galaxy: Galaxy) => void
  onDelete: (galaxy: Galaxy) => void
}

export function getGalaxyColumns({ onEdit, onDelete }: GalaxyColumnsProps): ColumnDef<Galaxy>[] {
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
      accessorKey: 'ownerUserID',
      header: 'Owner ID',
      cell: ({ row }) => {
        const id = row.getValue('ownerUserID') as string
        return (
          <span className="text-muted-foreground font-mono text-xs">
            {id.slice(0, 8)}...
          </span>
        )
      },
    },
    {
      accessorKey: 'enabled',
      header: 'Status',
      cell: ({ row }) => {
        const enabled = row.getValue('enabled') as boolean
        return (
          <Badge variant={enabled ? 'success' : 'secondary'}>
            {enabled ? 'Active' : 'Disabled'}
          </Badge>
        )
      },
    },
    {
      accessorKey: 'dateCreated',
      header: ({ column }) => <DataTableColumnHeader column={column} title="Created" />,
      cell: ({ row }) => {
        const date = new Date(row.getValue('dateCreated'))
        return <span className="text-muted-foreground">{date.toLocaleDateString()}</span>
      },
    },
    {
      id: 'actions',
      cell: ({ row }) => {
        const galaxy = row.original

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
              <DropdownMenuSeparator />
              <DropdownMenuItem onClick={() => onEdit(galaxy)}>
                <Pencil className="mr-2 h-4 w-4" />
                Edit
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={() => onDelete(galaxy)}
                className="text-destructive focus:text-destructive"
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
