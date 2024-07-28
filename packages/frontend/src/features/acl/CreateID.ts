/**
 * Generate UUID
 */
export const CreateID = (): string => {
  return crypto.randomUUID();
};
