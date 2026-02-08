import { z } from 'zod'

export const userRoleSchema = z.enum(['ADMIN', 'USER'])

export const createUserSchema = z
  .object({
    name: z.string().min(2, 'Name must be at least 2 characters'),
    email: z.string().email('Invalid email address'),
    password: z.string().min(8, 'Password must be at least 8 characters'),
    passwordConfirm: z.string(),
    roles: z.array(userRoleSchema).min(1, 'At least one role is required'),
    guild: z.string().optional(),
  })
  .refine((data) => data.password === data.passwordConfirm, {
    message: "Passwords don't match",
    path: ['passwordConfirm'],
  })

export const updateUserSchema = z.object({
  name: z.string().min(2, 'Name must be at least 2 characters').optional(),
  email: z.string().email('Invalid email address').optional(),
  roles: z.array(userRoleSchema).min(1, 'At least one role is required').optional(),
  guild: z.string().optional(),
  enabled: z.boolean().optional(),
})

export type CreateUserFormData = z.infer<typeof createUserSchema>
export type UpdateUserFormData = z.infer<typeof updateUserSchema>
