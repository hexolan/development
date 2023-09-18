import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import { panelsAdapter } from '../features/panels'
import { convertRawPanel } from '../types/panels'

import type { Panel } from '../types/common'
import type { GetPanelByNameRequest, CreatePanelRequest, UpdatePanelRequest, DeletePanelRequest, RawPanelResponse } from '../types/panels'

export const panelsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelByName: builder.query<EntityState<Panel>, GetPanelByNameRequest>({
      query: req => ({ url: `/v1/panels/${req.name}` }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    createPanel: builder.mutation<EntityState<Panel>, CreatePanelRequest>({
      query: req => ({
        url: '/v1/panels',
        method: 'POST',
        body: { ...req }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    updatePanel: builder.mutation<EntityState<Panel>, UpdatePanelRequest>({
      query: req => ({
        url: `/v1/panels/${req.name}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    deletePanel: builder.mutation<void, DeletePanelRequest>({
      query: req => ({
        url: `/v1/panels/${req.name}`,
        method: 'DELETE'
      })
    })
  })
})

export const { useGetPanelByNameQuery, useCreatePanelMutation, useUpdatePanelMutation, useDeletePanelMutation } = panelsApiSlice