import { z } from 'zod'

// Use z.number() directly without coercion - let form handle number conversion
const statField = z.number().min(0).max(1000)

export const createResourceSchema = z.object({
  name: z.string().min(1, 'Name is required'),
  galaxyID: z.string().uuid('Invalid galaxy ID'),
  addedUserID: z.string().uuid('Invalid user ID'),
  resourceType: z.string().min(1, 'Resource type is required'),
  // Stats
  cr: statField,
  cd: statField,
  dr: statField,
  fl: statField,
  hr: statField,
  ma: statField,
  pe: statField,
  oq: statField,
  sr: statField,
  ut: statField,
  er: statField,
})

export const updateResourceSchema = z.object({
  name: z.string().min(1, 'Name is required').optional(),
  galaxyID: z.string().uuid('Invalid galaxy ID').optional(),
  resourceType: z.string().min(1, 'Resource type is required').optional(),
  verified: z.boolean().optional(),
  // Stats - optional for updates
  cr: statField.optional(),
  cd: statField.optional(),
  dr: statField.optional(),
  fl: statField.optional(),
  hr: statField.optional(),
  ma: statField.optional(),
  pe: statField.optional(),
  oq: statField.optional(),
  sr: statField.optional(),
  ut: statField.optional(),
  er: statField.optional(),
})

export type CreateResourceFormData = z.infer<typeof createResourceSchema>
export type UpdateResourceFormData = z.infer<typeof updateResourceSchema>
