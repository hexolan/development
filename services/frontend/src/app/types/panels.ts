export type RawPanel = {
  id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at?: string;
}

export type CreatePanelData = {
  name: string;
  description: string;
}

export type UpdatePanelData = {
  name?: string;
  description?: string;
}