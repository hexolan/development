import { apiSlice } from '../api'

import { CreatePanelData, UpdatePanelData } from '../types/panels'

export const panelsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelByName: builder.query({
      query: (name: string) => ({
        url: `/v1/panels/${name}`
      })
    }),

    createPanel: builder.mutation({
      query: (data: CreatePanelData) => ({
        url: '/v1/panels',
        method: 'POST',
        body: { ...data }
      })
    }),

    updatePanel: builder.mutation({
      query: (data: UpdatePanelData) => ({
        url: `/v1/panels/${data.name}`,
        method: 'PATCH',
        body: { ...data }
      })
    }),

    deletePanel: builder.mutation({
      query: (name: string) => ({
        url: `/v1/panels/${name}`,
        method: 'DELETE'
      })
    })
  })
})

export const { useGetPanelByNameQuery, useCreatePanelMutation, useUpdatePanelMutation, useDeletePanelMutation } = panelsApiSlice