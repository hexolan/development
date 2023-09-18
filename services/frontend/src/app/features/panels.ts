import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Panel } from '../types/common'

export const panelsAdapter = createEntityAdapter<Panel>({
  selectId: (panel) => panel.id
})