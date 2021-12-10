import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const marketplaceApi = createApi({
  reducerPath: "marketplaceApi",
  baseQuery: fetchBaseQuery({
    baseUrl: "http://localhost:9000/",
    prepareHeaders: (headers, { getState }) => {
      const token = getState().user.token;
      if (token) {
        headers.set("authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    getUserById: builder.query({
      query: (id) => `users/${id}`,
    }),
    getUserBalance: builder.query({
      query: () => `users/balance`,
    }),
    getUserProducts: builder.query({
      query: (id) => `users/${id}/products`,
    }),
    getUserPurchasedProducts: builder.query({
      query: () => `users/purchased_products`,
    }),
    getUserSoldProducts: builder.query({
      query: () => `users/sold_products`,
    }),
    getUserTransactions: builder.query({
      query: () => `users/report/transactions`,
    }),
    getUserOrders: builder.query({
      query: () => `users/report/orders`,
    }),
    editUser: builder.mutation({
      query: (user) => ({
        url: `users`,
        method: "PUT",
        body: user,
      }),
    }),
    signIn: builder.mutation({
      query: (user) => ({
        url: `users/login`,
        method: "POST",
        body: user,
      }),
    }),
    signUp: builder.mutation({
      query: (user) => ({
        url: `users/signup`,
        method: "POST",
        body: user,
      }),
    }),
    depositMoney: builder.mutation({
      query: (amount) => ({
        url: `users/balance`,
        method: "POST",
        body: { amount },
      }),
    }),
    getAllProducts: builder.query({
      query: () => `products`,
    }),
    getProductById: builder.query({
      query: (id) => `products/${id}`,
    }),
    search: builder.query({
      query: (query) => `products/search?q=${query}`,
    }),
    addProduct: builder.mutation({
      query: (product) => ({
        url: `products`,
        method: "POST",
        body: product,
      }),
    }),
    orderProduct: builder.mutation({
      query: (id) => ({
        url: `products/${id}/order`,
        method: "POST",
      }),
    }),
    addProductToStore: builder.mutation({
      query: (id) => ({
        url: `products/${id}/store`,
        method: "POST",
      }),
    }),
    editProduct: builder.mutation({
      query: ({ id, product }) => ({
        url: `products/${id}`,
        method: "PUT",
        body: product,
      }),
    }),
    deleteProduct: builder.mutation({
      query: (id) => ({
        url: `products/${id}`,
        method: "DELETE",
      }),
    }),
    getStores: builder.query({
      query: () => `stores`,
    }),
    getStoreById: builder.query({
      query: (id) => `stores/${id}`,
    }),
  }),
});

export const {
  useGetUserByIdQuery,
  useGetUserBalanceQuery,
  useGetUserProductsQuery,
  useGetUserPurchasedProductsQuery,
  useGetUserSoldProductsQuery,
  useGetUserTransactionsQuery,
  useGetUserOrdersQuery,
  useEditUserMutation,
  useSignInMutation,
  useSignUpMutation,
  useDepositMoneyMutation,
  useGetAllProductsQuery,
  useGetProductByIdQuery,
  useLazySearchQuery,
  useAddProductMutation,
  useOrderProductMutation,
  useAddProductToStoreMutation,
  useEditProductMutation,
  useDeleteProductMutation,
  useGetStoresQuery,
  useGetStoreByIdQuery,
} = marketplaceApi;
