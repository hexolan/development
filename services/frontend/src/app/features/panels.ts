import { createEntityAdapter } from '@reduxjs/toolkit'

import type { Panel } from '../types/common'

const panelsAdapter = createEntityAdapter<Panel>({
  selectId: (panel) => panel.id
})

export default panelsAdapter