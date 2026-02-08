import { z } from 'zod'

export const createGalaxySchema = z.object({
  name: z.string().min(2, 'Name must be at least 2 characters'),
  ownerUserID: z.string().uuid('Invalid user ID'),
})

export const updateGalaxySchema = z.object({
  name: z.string().min(2, 'Name must be at least 2 characters').optional(),
  ownerUserID: z.string().uuid('Invalid user ID').optional(),
  enabled: z.boolean().optional(),
})

export type CreateGalaxyFormData = z.infer<typeof createGalaxySchema>
export type UpdateGalaxyFormData = z.infer<typeof updateGalaxySchema>
