import { convertRawTimestamp } from './api';

import type { Panel } from './common';
import type { RawTimestamp } from './api';

export const convertRawPanel = (rawPanel: RawPanel): Panel => ({
  id: rawPanel.id,
  name: rawPanel.name,
  description: rawPanel.description,
  createdAt: convertRawTimestamp(rawPanel.created_at),
  updatedAt: (rawPanel.updated_at ? convertRawTimestamp(rawPanel.updated_at) : undefined),
})

export type RawPanel = {
  id: string;
  name: string;
  description: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type CreatePanelData = {
  name: string;
  description: string;
}

export type UpdatePanelData = {
  name?: string;
  description?: string;
}