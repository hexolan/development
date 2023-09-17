import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Panel } from '../types'

const panelsAdapter = createEntityAdapter<Panel>({
  selectId: (panel) => panel.id
})

export default panelsAdapter