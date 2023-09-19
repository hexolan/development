import type { EntityState } from '@reduxjs/toolkit'

import { apiSlice } from '../api'
import { panelsAdapter } from '../features/panels'
import { convertRawPanel } from '../types/panels'

import type { Panel } from '../types/common'
import type {
  RawPanelResponse,
  GetPanelByIdRequest, GetPanelByNameRequest,
  UpdatePanelByIdRequest, UpdatePanelByNameRequest,
  DeletePanelByIdRequest, DeletePanelByNameRequest,
  CreatePanelRequest
} from '../types/panels'

// todo: create type GetPanelByIdRequest
// todo: seperate deletePanel into deleteById and updateByName
// todo: seperate updatePanel into updateById and deleteByName
// todo: rename any references to the mutations/query in other files
export const panelsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelById: builder.query<EntityState<Panel>, GetPanelByIdRequest>({
      query: req => ({ url: `/v1/panels/id/${req.id}` }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    getPanelByName: builder.query<EntityState<Panel>, GetPanelByNameRequest>({
      query: req => ({ url: `/v1/panels/name/${req.name}` }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    updatePanelById: builder.mutation<EntityState<Panel>, UpdatePanelByIdRequest>({
      query: req => ({
        url: `/v1/panels/id/${req.id}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    updatePanelByName: builder.mutation<EntityState<Panel>, UpdatePanelByNameRequest>({
      query: req => ({
        url: `/v1/panels/name/${req.name}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return panelsAdapter.setOne(panelsAdapter.getInitialState(), convertRawPanel(response.data))
      }
    }),

    deletePanelById: builder.mutation<void, DeletePanelByIdRequest>({
      query: req => ({
        url: `/v1/panels/id/${req.id}`,
        method: 'DELETE'
      })
    }),

    deletePanelByName: builder.mutation<void, DeletePanelByNameRequest>({
      query: req => ({
        url: `/v1/panels/id/${req.name}`,
        method: 'DELETE'
      })
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
  })
})

// todo: redo these:
// export const { useGetPanelByIdQuery, useGetPanelByNameQuery, useUpdatePanelMutation, useDeletePanelMutation, useCreatePanelMutation } = panelsApiSlice