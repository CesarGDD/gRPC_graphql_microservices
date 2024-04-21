/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
import gql from 'graphql-tag';
import * as Urql from 'urql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type AddItemInput = {
  item: ProductItemInput;
  userId: Scalars['Int']['input'];
};

export type CartResponse = {
  __typename?: 'CartResponse';
  cart: ShoppingCart;
  success: Scalars['Boolean']['output'];
};

export type ClearCartInput = {
  userId: Scalars['Int']['input'];
};

export type CreateOrderInput = {
  orderItems: Array<NewOrderItem>;
  totalPrice: Scalars['Int']['input'];
  userId: Scalars['Int']['input'];
};

export type GetCartInput = {
  userId: Scalars['Int']['input'];
};

export type GetOrderDetailsInput = {
  orderId: Scalars['Int']['input'];
};

export type GetOrdersDetailsByUserIdInput = {
  userId: Scalars['Int']['input'];
};

export type GetPaymentDetailsInput = {
  orderId: Scalars['Int']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addItem: Response;
  clearCart: Response;
  createOrder: Response;
  createProduct: Response;
  deleteProduct: Response;
  deleteUser: Response;
  processPayment: Response;
  registerUser: Response;
  removeItem: Response;
  updateItem: Response;
  updateProduct: Response;
  updateUser: Response;
};


export type MutationAddItemArgs = {
  input: AddItemInput;
};


export type MutationClearCartArgs = {
  input: ClearCartInput;
};


export type MutationCreateOrderArgs = {
  input: CreateOrderInput;
};


export type MutationCreateProductArgs = {
  input: NewProductInput;
};


export type MutationDeleteProductArgs = {
  productId: Scalars['Int']['input'];
};


export type MutationDeleteUserArgs = {
  userId: Scalars['Int']['input'];
};


export type MutationProcessPaymentArgs = {
  input: ProcessPaymentInput;
};


export type MutationRegisterUserArgs = {
  input: RegisterUserInput;
};


export type MutationRemoveItemArgs = {
  input: RemoveItemInput;
};


export type MutationUpdateItemArgs = {
  input: UpdateItemInput;
};


export type MutationUpdateProductArgs = {
  input: UpdateProductInput;
};


export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};

export type NewOrderItem = {
  productId: Scalars['Int']['input'];
  quantity: Scalars['Int']['input'];
};

export type NewProductInput = {
  description: Scalars['String']['input'];
  name: Scalars['String']['input'];
  price: Scalars['Int']['input'];
  title: Scalars['String']['input'];
  url: Scalars['String']['input'];
  userId: Scalars['Int']['input'];
};

export type Order = {
  __typename?: 'Order';
  orderDetails: OrderDetails;
};

export type OrderDetails = {
  __typename?: 'OrderDetails';
  orderId: Scalars['Int']['output'];
  orderItems: Array<OrderItem>;
  status: Scalars['String']['output'];
  totalPrice: Scalars['Int']['output'];
  userId: Scalars['Int']['output'];
};

export type OrderItem = {
  __typename?: 'OrderItem';
  productId: Scalars['Int']['output'];
  quantity: Scalars['Int']['output'];
};

export type PaymentDetails = {
  __typename?: 'PaymentDetails';
  paymentMethod: Scalars['String']['output'];
};

export type ProcessPaymentInput = {
  orderId: Scalars['Int']['input'];
  paymentMethod: Scalars['String']['input'];
};

export type Product = {
  __typename?: 'Product';
  description: Scalars['String']['output'];
  name: Scalars['String']['output'];
  price: Scalars['Int']['output'];
  productId: Scalars['Int']['output'];
  title: Scalars['String']['output'];
  url: Scalars['String']['output'];
  userId: Scalars['Int']['output'];
  userOwner: User;
};

export type ProductItem = {
  __typename?: 'ProductItem';
  product: Product;
  productId: Scalars['Int']['output'];
  quantity: Scalars['Int']['output'];
};

export type ProductItemInput = {
  productId: Scalars['Int']['input'];
  quantity: Scalars['Int']['input'];
};

export type ProductResponse = {
  __typename?: 'ProductResponse';
  product: Product;
};

export type Query = {
  __typename?: 'Query';
  GetOrdersDetailsByUserId: Array<OrderDetails>;
  getCart: CartResponse;
  getOrderDetails: Order;
  getPaymentDetails: PaymentDetails;
  getProduct: ProductResponse;
  getProductByName: ProductResponse;
  getProducts: Array<ProductResponse>;
  getUser: UserResponse;
  getUserByUsername: UserResponse;
  getUsers: Array<UserResponse>;
};


export type QueryGetOrdersDetailsByUserIdArgs = {
  input: GetOrdersDetailsByUserIdInput;
};


export type QueryGetCartArgs = {
  input: GetCartInput;
};


export type QueryGetOrderDetailsArgs = {
  input: GetOrderDetailsInput;
};


export type QueryGetPaymentDetailsArgs = {
  input: GetPaymentDetailsInput;
};


export type QueryGetProductArgs = {
  productId: Scalars['Int']['input'];
};


export type QueryGetProductByNameArgs = {
  name: Scalars['String']['input'];
};


export type QueryGetUserArgs = {
  userId: Scalars['Int']['input'];
};


export type QueryGetUserByUsernameArgs = {
  username: Scalars['String']['input'];
};

export type RegisterUserInput = {
  password: Scalars['String']['input'];
  role: Role;
  username: Scalars['String']['input'];
};

export type RemoveItemInput = {
  productId: Scalars['Int']['input'];
  userId: Scalars['Int']['input'];
};

export type Response = {
  __typename?: 'Response';
  message: Scalars['String']['output'];
  success: Scalars['Boolean']['output'];
};

export enum Role {
  Admin = 'ADMIN',
  Customer = 'CUSTOMER'
}

export type ShoppingCart = {
  __typename?: 'ShoppingCart';
  items: Array<ProductItem>;
  userId: Scalars['Int']['output'];
};

export type UpdateItemInput = {
  item: ProductItemInput;
  userId: Scalars['Int']['input'];
};

export type UpdateProductInput = {
  description: Scalars['String']['input'];
  name: Scalars['String']['input'];
  price: Scalars['Int']['input'];
  productId: Scalars['Int']['input'];
  title: Scalars['String']['input'];
  url: Scalars['String']['input'];
};

export type UpdateUserInput = {
  password: Scalars['String']['input'];
  role: Role;
  userId: Scalars['Int']['input'];
  username: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  orders?: Maybe<Array<OrderDetails>>;
  products?: Maybe<Array<Product>>;
  role: Role;
  userId: Scalars['Int']['output'];
  username: Scalars['String']['output'];
};

export type UserResponse = {
  __typename?: 'UserResponse';
  user: User;
};

export type GetOrdersDetailsByUserIdQueryVariables = Exact<{
  input: GetOrdersDetailsByUserIdInput;
}>;


export type GetOrdersDetailsByUserIdQuery = { __typename?: 'Query', GetOrdersDetailsByUserId: Array<{ __typename?: 'OrderDetails', orderId: number, status: string, totalPrice: number, userId: number, orderItems: Array<{ __typename?: 'OrderItem', productId: number, quantity: number }> }> };

export type GetOrderDetailsQueryVariables = Exact<{
  input: GetOrderDetailsInput;
}>;


export type GetOrderDetailsQuery = { __typename?: 'Query', getOrderDetails: { __typename?: 'Order', orderDetails: { __typename?: 'OrderDetails', orderId: number, status: string, totalPrice: number, userId: number, orderItems: Array<{ __typename?: 'OrderItem', productId: number, quantity: number }> } } };

export type GetPaymentDetailsQueryVariables = Exact<{
  input: GetPaymentDetailsInput;
}>;


export type GetPaymentDetailsQuery = { __typename?: 'Query', getPaymentDetails: { __typename?: 'PaymentDetails', paymentMethod: string } };

export type CreateOrderMutationVariables = Exact<{
  input: CreateOrderInput;
}>;


export type CreateOrderMutation = { __typename?: 'Mutation', createOrder: { __typename?: 'Response', message: string, success: boolean } };

export type ProcessPaymentMutationVariables = Exact<{
  input: ProcessPaymentInput;
}>;


export type ProcessPaymentMutation = { __typename?: 'Mutation', processPayment: { __typename?: 'Response', message: string, success: boolean } };

export type GetProductsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetProductsQuery = { __typename?: 'Query', getProducts: Array<{ __typename?: 'ProductResponse', product: { __typename?: 'Product', description: string, name: string, price: number, productId: number, title: string, url: string, userId: number, userOwner: { __typename?: 'User', username: string, userId: number } } }> };

export type GetProductByNameQueryVariables = Exact<{
  name: Scalars['String']['input'];
}>;


export type GetProductByNameQuery = { __typename?: 'Query', getProductByName: { __typename?: 'ProductResponse', product: { __typename?: 'Product', description: string, name: string, price: number, productId: number, title: string, url: string, userId: number, userOwner: { __typename?: 'User', userId: number, username: string } } } };

export type GetProductQueryVariables = Exact<{
  productId: Scalars['Int']['input'];
}>;


export type GetProductQuery = { __typename?: 'Query', getProduct: { __typename?: 'ProductResponse', product: { __typename?: 'Product', description: string, name: string, price: number, productId: number, title: string, url: string, userId: number, userOwner: { __typename?: 'User', userId: number, username: string } } } };

export type CreateProductMutationVariables = Exact<{
  input: NewProductInput;
}>;


export type CreateProductMutation = { __typename?: 'Mutation', createProduct: { __typename?: 'Response', message: string, success: boolean } };

export type DeleteProductMutationVariables = Exact<{
  productId: Scalars['Int']['input'];
}>;


export type DeleteProductMutation = { __typename?: 'Mutation', deleteProduct: { __typename?: 'Response', message: string, success: boolean } };

export type UpdateProductMutationVariables = Exact<{
  input: UpdateProductInput;
}>;


export type UpdateProductMutation = { __typename?: 'Mutation', updateProduct: { __typename?: 'Response', message: string, success: boolean } };

export type GetCartQueryVariables = Exact<{
  input: GetCartInput;
}>;


export type GetCartQuery = { __typename?: 'Query', getCart: { __typename?: 'CartResponse', cart: { __typename?: 'ShoppingCart', items: Array<{ __typename?: 'ProductItem', productId: number, quantity: number, product: { __typename?: 'Product', name: string, price: number, url: string, userOwner: { __typename?: 'User', username: string, userId: number } } }> } } };

export type AddItemMutationVariables = Exact<{
  input: AddItemInput;
}>;


export type AddItemMutation = { __typename?: 'Mutation', addItem: { __typename?: 'Response', message: string, success: boolean } };

export type ClearCartMutationVariables = Exact<{
  input: ClearCartInput;
}>;


export type ClearCartMutation = { __typename?: 'Mutation', clearCart: { __typename?: 'Response', message: string, success: boolean } };

export type RemoveItemMutationVariables = Exact<{
  input: RemoveItemInput;
}>;


export type RemoveItemMutation = { __typename?: 'Mutation', removeItem: { __typename?: 'Response', message: string, success: boolean } };

export type UpdateItemMutationVariables = Exact<{
  input: UpdateItemInput;
}>;


export type UpdateItemMutation = { __typename?: 'Mutation', updateItem: { __typename?: 'Response', message: string, success: boolean } };

export type GetUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUsersQuery = { __typename?: 'Query', getUsers: Array<{ __typename?: 'UserResponse', user: { __typename?: 'User', username: string, userId: number, role: Role } }> };

export type GetUserQueryVariables = Exact<{
  userId: Scalars['Int']['input'];
}>;


export type GetUserQuery = { __typename?: 'Query', getUser: { __typename?: 'UserResponse', user: { __typename?: 'User', username: string, userId: number, role: Role } } };

export type GetUserByUsernameQueryVariables = Exact<{
  username: Scalars['String']['input'];
}>;


export type GetUserByUsernameQuery = { __typename?: 'Query', getUserByUsername: { __typename?: 'UserResponse', user: { __typename?: 'User', username: string, userId: number, role: Role } } };

export type RegisterUserMutationVariables = Exact<{
  input: RegisterUserInput;
}>;


export type RegisterUserMutation = { __typename?: 'Mutation', registerUser: { __typename?: 'Response', message: string, success: boolean } };

export type UpdateUserMutationVariables = Exact<{
  input: UpdateUserInput;
}>;


export type UpdateUserMutation = { __typename?: 'Mutation', updateUser: { __typename?: 'Response', message: string, success: boolean } };

export type DeleteUserMutationVariables = Exact<{
  userId: Scalars['Int']['input'];
}>;


export type DeleteUserMutation = { __typename?: 'Mutation', deleteUser: { __typename?: 'Response', message: string, success: boolean } };


export const GetOrdersDetailsByUserIdDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetOrdersDetailsByUserId"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"GetOrdersDetailsByUserIdInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"GetOrdersDetailsByUserId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orderId"}},{"kind":"Field","name":{"kind":"Name","value":"orderItems"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productId"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"totalPrice"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}}]}}]}}]} as unknown as DocumentNode<GetOrdersDetailsByUserIdQuery, GetOrdersDetailsByUserIdQueryVariables>;
export const GetOrderDetailsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getOrderDetails"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"GetOrderDetailsInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getOrderDetails"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orderDetails"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orderId"}},{"kind":"Field","name":{"kind":"Name","value":"orderItems"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productId"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"totalPrice"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}}]}}]}}]}}]} as unknown as DocumentNode<GetOrderDetailsQuery, GetOrderDetailsQueryVariables>;
export const GetPaymentDetailsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getPaymentDetails"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"GetPaymentDetailsInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getPaymentDetails"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"paymentMethod"}}]}}]}}]} as unknown as DocumentNode<GetPaymentDetailsQuery, GetPaymentDetailsQueryVariables>;
export const CreateOrderDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrder"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateOrderInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrder"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<CreateOrderMutation, CreateOrderMutationVariables>;
export const ProcessPaymentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"processPayment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ProcessPaymentInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"processPayment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<ProcessPaymentMutation, ProcessPaymentMutationVariables>;
export const GetProductsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetProducts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getProducts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"price"}},{"kind":"Field","name":{"kind":"Name","value":"productId"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"url"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"userOwner"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}}]}}]}}]}}]}}]} as unknown as DocumentNode<GetProductsQuery, GetProductsQueryVariables>;
export const GetProductByNameDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getProductByName"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getProductByName"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"price"}},{"kind":"Field","name":{"kind":"Name","value":"productId"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"url"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"userOwner"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}}]}}]}}]}}]} as unknown as DocumentNode<GetProductByNameQuery, GetProductByNameQueryVariables>;
export const GetProductDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetProduct"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"productId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getProduct"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"productId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"productId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"price"}},{"kind":"Field","name":{"kind":"Name","value":"productId"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"url"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"userOwner"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}}]}}]}}]}}]} as unknown as DocumentNode<GetProductQuery, GetProductQueryVariables>;
export const CreateProductDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateProduct"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewProductInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createProduct"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<CreateProductMutation, CreateProductMutationVariables>;
export const DeleteProductDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"DeleteProduct"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"productId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteProduct"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"productId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"productId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<DeleteProductMutation, DeleteProductMutationVariables>;
export const UpdateProductDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdateProduct"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UpdateProductInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateProduct"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<UpdateProductMutation, UpdateProductMutationVariables>;
export const GetCartDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetCart"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"GetCartInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getCart"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cart"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productId"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"price"}},{"kind":"Field","name":{"kind":"Name","value":"url"}},{"kind":"Field","name":{"kind":"Name","value":"userOwner"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<GetCartQuery, GetCartQueryVariables>;
export const AddItemDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"AddItem"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"AddItemInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addItem"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<AddItemMutation, AddItemMutationVariables>;
export const ClearCartDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"ClearCart"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ClearCartInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"clearCart"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<ClearCartMutation, ClearCartMutationVariables>;
export const RemoveItemDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"RemoveItem"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RemoveItemInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"removeItem"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<RemoveItemMutation, RemoveItemMutationVariables>;
export const UpdateItemDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdateItem"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UpdateItemInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateItem"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<UpdateItemMutation, UpdateItemMutationVariables>;
export const GetUsersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetUsers"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUsers"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}}]}}]}}]} as unknown as DocumentNode<GetUsersQuery, GetUsersQueryVariables>;
export const GetUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}}]}}]}}]} as unknown as DocumentNode<GetUserQuery, GetUserQueryVariables>;
export const GetUserByUsernameDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetUserByUsername"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"username"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUserByUsername"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"username"},"value":{"kind":"Variable","name":{"kind":"Name","value":"username"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}}]}}]}}]} as unknown as DocumentNode<GetUserByUsernameQuery, GetUserByUsernameQueryVariables>;
export const RegisterUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"RegisterUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RegisterUserInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"registerUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<RegisterUserMutation, RegisterUserMutationVariables>;
export const UpdateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"UpdateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UpdateUserInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<UpdateUserMutation, UpdateUserMutationVariables>;
export const DeleteUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"DeleteUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<DeleteUserMutation, DeleteUserMutationVariables>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type AddItemInput = {
  item: ProductItemInput;
  userId: Scalars['Int']['input'];
};

export type CartResponse = {
  __typename?: 'CartResponse';
  cart: ShoppingCart;
  success: Scalars['Boolean']['output'];
};

export type ClearCartInput = {
  userId: Scalars['Int']['input'];
};

export type CreateOrderInput = {
  orderItems: Array<NewOrderItem>;
  totalPrice: Scalars['Int']['input'];
  userId: Scalars['Int']['input'];
};

export type GetCartInput = {
  userId: Scalars['Int']['input'];
};

export type GetOrderDetailsInput = {
  orderId: Scalars['Int']['input'];
};

export type GetOrdersDetailsByUserIdInput = {
  userId: Scalars['Int']['input'];
};

export type GetPaymentDetailsInput = {
  orderId: Scalars['Int']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addItem: Response;
  clearCart: Response;
  createOrder: Response;
  createProduct: Response;
  deleteProduct: Response;
  deleteUser: Response;
  processPayment: Response;
  registerUser: Response;
  removeItem: Response;
  updateItem: Response;
  updateProduct: Response;
  updateUser: Response;
};


export type MutationAddItemArgs = {
  input: AddItemInput;
};


export type MutationClearCartArgs = {
  input: ClearCartInput;
};


export type MutationCreateOrderArgs = {
  input: CreateOrderInput;
};


export type MutationCreateProductArgs = {
  input: NewProductInput;
};


export type MutationDeleteProductArgs = {
  productId: Scalars['Int']['input'];
};


export type MutationDeleteUserArgs = {
  userId: Scalars['Int']['input'];
};


export type MutationProcessPaymentArgs = {
  input: ProcessPaymentInput;
};


export type MutationRegisterUserArgs = {
  input: RegisterUserInput;
};


export type MutationRemoveItemArgs = {
  input: RemoveItemInput;
};


export type MutationUpdateItemArgs = {
  input: UpdateItemInput;
};


export type MutationUpdateProductArgs = {
  input: UpdateProductInput;
};


export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};

export type NewOrderItem = {
  productId: Scalars['Int']['input'];
  quantity: Scalars['Int']['input'];
};

export type NewProductInput = {
  description: Scalars['String']['input'];
  name: Scalars['String']['input'];
  price: Scalars['Int']['input'];
  title: Scalars['String']['input'];
  url: Scalars['String']['input'];
  userId: Scalars['Int']['input'];
};

export type Order = {
  __typename?: 'Order';
  orderDetails: OrderDetails;
};

export type OrderDetails = {
  __typename?: 'OrderDetails';
  orderId: Scalars['Int']['output'];
  orderItems: Array<OrderItem>;
  status: Scalars['String']['output'];
  totalPrice: Scalars['Int']['output'];
  userId: Scalars['Int']['output'];
};

export type OrderItem = {
  __typename?: 'OrderItem';
  productId: Scalars['Int']['output'];
  quantity: Scalars['Int']['output'];
};

export type PaymentDetails = {
  __typename?: 'PaymentDetails';
  paymentMethod: Scalars['String']['output'];
};

export type ProcessPaymentInput = {
  orderId: Scalars['Int']['input'];
  paymentMethod: Scalars['String']['input'];
};

export type Product = {
  __typename?: 'Product';
  description: Scalars['String']['output'];
  name: Scalars['String']['output'];
  price: Scalars['Int']['output'];
  productId: Scalars['Int']['output'];
  title: Scalars['String']['output'];
  url: Scalars['String']['output'];
  userId: Scalars['Int']['output'];
  userOwner: User;
};

export type ProductItem = {
  __typename?: 'ProductItem';
  product: Product;
  productId: Scalars['Int']['output'];
  quantity: Scalars['Int']['output'];
};

export type ProductItemInput = {
  productId: Scalars['Int']['input'];
  quantity: Scalars['Int']['input'];
};

export type ProductResponse = {
  __typename?: 'ProductResponse';
  product: Product;
};

export type Query = {
  __typename?: 'Query';
  GetOrdersDetailsByUserId: Array<OrderDetails>;
  getCart: CartResponse;
  getOrderDetails: Order;
  getPaymentDetails: PaymentDetails;
  getProduct: ProductResponse;
  getProductByName: ProductResponse;
  getProducts: Array<ProductResponse>;
  getUser: UserResponse;
  getUserByUsername: UserResponse;
  getUsers: Array<UserResponse>;
};


export type QueryGetOrdersDetailsByUserIdArgs = {
  input: GetOrdersDetailsByUserIdInput;
};


export type QueryGetCartArgs = {
  input: GetCartInput;
};


export type QueryGetOrderDetailsArgs = {
  input: GetOrderDetailsInput;
};


export type QueryGetPaymentDetailsArgs = {
  input: GetPaymentDetailsInput;
};


export type QueryGetProductArgs = {
  productId: Scalars['Int']['input'];
};


export type QueryGetProductByNameArgs = {
  name: Scalars['String']['input'];
};


export type QueryGetUserArgs = {
  userId: Scalars['Int']['input'];
};


export type QueryGetUserByUsernameArgs = {
  username: Scalars['String']['input'];
};

export type RegisterUserInput = {
  password: Scalars['String']['input'];
  role: Role;
  username: Scalars['String']['input'];
};

export type RemoveItemInput = {
  productId: Scalars['Int']['input'];
  userId: Scalars['Int']['input'];
};

export type Response = {
  __typename?: 'Response';
  message: Scalars['String']['output'];
  success: Scalars['Boolean']['output'];
};

export enum Role {
  Admin = 'ADMIN',
  Customer = 'CUSTOMER'
}

export type ShoppingCart = {
  __typename?: 'ShoppingCart';
  items: Array<ProductItem>;
  userId: Scalars['Int']['output'];
};

export type UpdateItemInput = {
  item: ProductItemInput;
  userId: Scalars['Int']['input'];
};

export type UpdateProductInput = {
  description: Scalars['String']['input'];
  name: Scalars['String']['input'];
  price: Scalars['Int']['input'];
  productId: Scalars['Int']['input'];
  title: Scalars['String']['input'];
  url: Scalars['String']['input'];
};

export type UpdateUserInput = {
  password: Scalars['String']['input'];
  role: Role;
  userId: Scalars['Int']['input'];
  username: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  orders?: Maybe<Array<OrderDetails>>;
  products?: Maybe<Array<Product>>;
  role: Role;
  userId: Scalars['Int']['output'];
  username: Scalars['String']['output'];
};

export type UserResponse = {
  __typename?: 'UserResponse';
  user: User;
};

export type GetOrdersDetailsByUserIdQueryVariables = Exact<{
  input: GetOrdersDetailsByUserIdInput;
}>;


export type GetOrdersDetailsByUserIdQuery = { __typename?: 'Query', GetOrdersDetailsByUserId: Array<{ __typename?: 'OrderDetails', orderId: number, status: string, totalPrice: number, userId: number, orderItems: Array<{ __typename?: 'OrderItem', productId: number, quantity: number }> }> };

export type GetOrderDetailsQueryVariables = Exact<{
  input: GetOrderDetailsInput;
}>;


export type GetOrderDetailsQuery = { __typename?: 'Query', getOrderDetails: { __typename?: 'Order', orderDetails: { __typename?: 'OrderDetails', orderId: number, status: string, totalPrice: number, userId: number, orderItems: Array<{ __typename?: 'OrderItem', productId: number, quantity: number }> } } };

export type GetPaymentDetailsQueryVariables = Exact<{
  input: GetPaymentDetailsInput;
}>;


export type GetPaymentDetailsQuery = { __typename?: 'Query', getPaymentDetails: { __typename?: 'PaymentDetails', paymentMethod: string } };

export type CreateOrderMutationVariables = Exact<{
  input: CreateOrderInput;
}>;


export type CreateOrderMutation = { __typename?: 'Mutation', createOrder: { __typename?: 'Response', message: string, success: boolean } };

export type ProcessPaymentMutationVariables = Exact<{
  input: ProcessPaymentInput;
}>;


export type ProcessPaymentMutation = { __typename?: 'Mutation', processPayment: { __typename?: 'Response', message: string, success: boolean } };

export type GetProductsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetProductsQuery = { __typename?: 'Query', getProducts: Array<{ __typename?: 'ProductResponse', product: { __typename?: 'Product', description: string, name: string, price: number, productId: number, title: string, url: string, userId: number, userOwner: { __typename?: 'User', username: string, userId: number } } }> };

export type GetProductByNameQueryVariables = Exact<{
  name: Scalars['String']['input'];
}>;


export type GetProductByNameQuery = { __typename?: 'Query', getProductByName: { __typename?: 'ProductResponse', product: { __typename?: 'Product', description: string, name: string, price: number, productId: number, title: string, url: string, userId: number, userOwner: { __typename?: 'User', userId: number, username: string } } } };

export type GetProductQueryVariables = Exact<{
  productId: Scalars['Int']['input'];
}>;


export type GetProductQuery = { __typename?: 'Query', getProduct: { __typename?: 'ProductResponse', product: { __typename?: 'Product', description: string, name: string, price: number, productId: number, title: string, url: string, userId: number, userOwner: { __typename?: 'User', userId: number, username: string } } } };

export type CreateProductMutationVariables = Exact<{
  input: NewProductInput;
}>;


export type CreateProductMutation = { __typename?: 'Mutation', createProduct: { __typename?: 'Response', message: string, success: boolean } };

export type DeleteProductMutationVariables = Exact<{
  productId: Scalars['Int']['input'];
}>;


export type DeleteProductMutation = { __typename?: 'Mutation', deleteProduct: { __typename?: 'Response', message: string, success: boolean } };

export type UpdateProductMutationVariables = Exact<{
  input: UpdateProductInput;
}>;


export type UpdateProductMutation = { __typename?: 'Mutation', updateProduct: { __typename?: 'Response', message: string, success: boolean } };

export type GetCartQueryVariables = Exact<{
  input: GetCartInput;
}>;


export type GetCartQuery = { __typename?: 'Query', getCart: { __typename?: 'CartResponse', cart: { __typename?: 'ShoppingCart', items: Array<{ __typename?: 'ProductItem', productId: number, quantity: number, product: { __typename?: 'Product', name: string, price: number, url: string, userOwner: { __typename?: 'User', username: string, userId: number } } }> } } };

export type AddItemMutationVariables = Exact<{
  input: AddItemInput;
}>;


export type AddItemMutation = { __typename?: 'Mutation', addItem: { __typename?: 'Response', message: string, success: boolean } };

export type ClearCartMutationVariables = Exact<{
  input: ClearCartInput;
}>;


export type ClearCartMutation = { __typename?: 'Mutation', clearCart: { __typename?: 'Response', message: string, success: boolean } };

export type RemoveItemMutationVariables = Exact<{
  input: RemoveItemInput;
}>;


export type RemoveItemMutation = { __typename?: 'Mutation', removeItem: { __typename?: 'Response', message: string, success: boolean } };

export type UpdateItemMutationVariables = Exact<{
  input: UpdateItemInput;
}>;


export type UpdateItemMutation = { __typename?: 'Mutation', updateItem: { __typename?: 'Response', message: string, success: boolean } };

export type GetUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUsersQuery = { __typename?: 'Query', getUsers: Array<{ __typename?: 'UserResponse', user: { __typename?: 'User', username: string, userId: number, role: Role } }> };

export type GetUserQueryVariables = Exact<{
  userId: Scalars['Int']['input'];
}>;


export type GetUserQuery = { __typename?: 'Query', getUser: { __typename?: 'UserResponse', user: { __typename?: 'User', username: string, userId: number, role: Role } } };

export type GetUserByUsernameQueryVariables = Exact<{
  username: Scalars['String']['input'];
}>;


export type GetUserByUsernameQuery = { __typename?: 'Query', getUserByUsername: { __typename?: 'UserResponse', user: { __typename?: 'User', username: string, userId: number, role: Role } } };

export type RegisterUserMutationVariables = Exact<{
  input: RegisterUserInput;
}>;


export type RegisterUserMutation = { __typename?: 'Mutation', registerUser: { __typename?: 'Response', message: string, success: boolean } };

export type UpdateUserMutationVariables = Exact<{
  input: UpdateUserInput;
}>;


export type UpdateUserMutation = { __typename?: 'Mutation', updateUser: { __typename?: 'Response', message: string, success: boolean } };

export type DeleteUserMutationVariables = Exact<{
  userId: Scalars['Int']['input'];
}>;


export type DeleteUserMutation = { __typename?: 'Mutation', deleteUser: { __typename?: 'Response', message: string, success: boolean } };


export const GetOrdersDetailsByUserIdDocument = gql`
    query GetOrdersDetailsByUserId($input: GetOrdersDetailsByUserIdInput!) {
  GetOrdersDetailsByUserId(input: $input) {
    orderId
    orderItems {
      productId
      quantity
    }
    status
    totalPrice
    userId
  }
}
    `;

export function useGetOrdersDetailsByUserIdQuery(options: Omit<Urql.UseQueryArgs<GetOrdersDetailsByUserIdQueryVariables>, 'query'>) {
  return Urql.useQuery<GetOrdersDetailsByUserIdQuery, GetOrdersDetailsByUserIdQueryVariables>({ query: GetOrdersDetailsByUserIdDocument, ...options });
};
export const GetOrderDetailsDocument = gql`
    query getOrderDetails($input: GetOrderDetailsInput!) {
  getOrderDetails(input: $input) {
    orderDetails {
      orderId
      orderItems {
        productId
        quantity
      }
      status
      totalPrice
      userId
    }
  }
}
    `;

export function useGetOrderDetailsQuery(options: Omit<Urql.UseQueryArgs<GetOrderDetailsQueryVariables>, 'query'>) {
  return Urql.useQuery<GetOrderDetailsQuery, GetOrderDetailsQueryVariables>({ query: GetOrderDetailsDocument, ...options });
};
export const GetPaymentDetailsDocument = gql`
    query getPaymentDetails($input: GetPaymentDetailsInput!) {
  getPaymentDetails(input: $input) {
    paymentMethod
  }
}
    `;

export function useGetPaymentDetailsQuery(options: Omit<Urql.UseQueryArgs<GetPaymentDetailsQueryVariables>, 'query'>) {
  return Urql.useQuery<GetPaymentDetailsQuery, GetPaymentDetailsQueryVariables>({ query: GetPaymentDetailsDocument, ...options });
};
export const CreateOrderDocument = gql`
    mutation createOrder($input: CreateOrderInput!) {
  createOrder(input: $input) {
    message
    success
  }
}
    `;

export function useCreateOrderMutation() {
  return Urql.useMutation<CreateOrderMutation, CreateOrderMutationVariables>(CreateOrderDocument);
};
export const ProcessPaymentDocument = gql`
    mutation processPayment($input: ProcessPaymentInput!) {
  processPayment(input: $input) {
    message
    success
  }
}
    `;

export function useProcessPaymentMutation() {
  return Urql.useMutation<ProcessPaymentMutation, ProcessPaymentMutationVariables>(ProcessPaymentDocument);
};
export const GetProductsDocument = gql`
    query GetProducts {
  getProducts {
    product {
      description
      name
      price
      productId
      title
      url
      userId
      userOwner {
        username
        userId
      }
    }
  }
}
    `;

export function useGetProductsQuery(options?: Omit<Urql.UseQueryArgs<GetProductsQueryVariables>, 'query'>) {
  return Urql.useQuery<GetProductsQuery, GetProductsQueryVariables>({ query: GetProductsDocument, ...options });
};
export const GetProductByNameDocument = gql`
    query getProductByName($name: String!) {
  getProductByName(name: $name) {
    product {
      description
      name
      price
      productId
      title
      url
      userId
      userOwner {
        userId
        username
      }
    }
  }
}
    `;

export function useGetProductByNameQuery(options: Omit<Urql.UseQueryArgs<GetProductByNameQueryVariables>, 'query'>) {
  return Urql.useQuery<GetProductByNameQuery, GetProductByNameQueryVariables>({ query: GetProductByNameDocument, ...options });
};
export const GetProductDocument = gql`
    query GetProduct($productId: Int!) {
  getProduct(productId: $productId) {
    product {
      description
      name
      price
      productId
      title
      url
      userId
      userOwner {
        userId
        username
      }
    }
  }
}
    `;

export function useGetProductQuery(options: Omit<Urql.UseQueryArgs<GetProductQueryVariables>, 'query'>) {
  return Urql.useQuery<GetProductQuery, GetProductQueryVariables>({ query: GetProductDocument, ...options });
};
export const CreateProductDocument = gql`
    mutation CreateProduct($input: NewProductInput!) {
  createProduct(input: $input) {
    message
    success
  }
}
    `;

export function useCreateProductMutation() {
  return Urql.useMutation<CreateProductMutation, CreateProductMutationVariables>(CreateProductDocument);
};
export const DeleteProductDocument = gql`
    mutation DeleteProduct($productId: Int!) {
  deleteProduct(productId: $productId) {
    message
    success
  }
}
    `;

export function useDeleteProductMutation() {
  return Urql.useMutation<DeleteProductMutation, DeleteProductMutationVariables>(DeleteProductDocument);
};
export const UpdateProductDocument = gql`
    mutation UpdateProduct($input: UpdateProductInput!) {
  updateProduct(input: $input) {
    message
    success
  }
}
    `;

export function useUpdateProductMutation() {
  return Urql.useMutation<UpdateProductMutation, UpdateProductMutationVariables>(UpdateProductDocument);
};
export const GetCartDocument = gql`
    query GetCart($input: GetCartInput!) {
  getCart(input: $input) {
    cart {
      items {
        productId
        quantity
        product {
          name
          price
          url
          userOwner {
            username
            userId
          }
        }
      }
    }
  }
}
    `;

export function useGetCartQuery(options: Omit<Urql.UseQueryArgs<GetCartQueryVariables>, 'query'>) {
  return Urql.useQuery<GetCartQuery, GetCartQueryVariables>({ query: GetCartDocument, ...options });
};
export const AddItemDocument = gql`
    mutation AddItem($input: AddItemInput!) {
  addItem(input: $input) {
    message
    success
  }
}
    `;

export function useAddItemMutation() {
  return Urql.useMutation<AddItemMutation, AddItemMutationVariables>(AddItemDocument);
};
export const ClearCartDocument = gql`
    mutation ClearCart($input: ClearCartInput!) {
  clearCart(input: $input) {
    message
    success
  }
}
    `;

export function useClearCartMutation() {
  return Urql.useMutation<ClearCartMutation, ClearCartMutationVariables>(ClearCartDocument);
};
export const RemoveItemDocument = gql`
    mutation RemoveItem($input: RemoveItemInput!) {
  removeItem(input: $input) {
    message
    success
  }
}
    `;

export function useRemoveItemMutation() {
  return Urql.useMutation<RemoveItemMutation, RemoveItemMutationVariables>(RemoveItemDocument);
};
export const UpdateItemDocument = gql`
    mutation UpdateItem($input: UpdateItemInput!) {
  updateItem(input: $input) {
    message
    success
  }
}
    `;

export function useUpdateItemMutation() {
  return Urql.useMutation<UpdateItemMutation, UpdateItemMutationVariables>(UpdateItemDocument);
};
export const GetUsersDocument = gql`
    query GetUsers {
  getUsers {
    user {
      username
      userId
      role
    }
  }
}
    `;

export function useGetUsersQuery(options?: Omit<Urql.UseQueryArgs<GetUsersQueryVariables>, 'query'>) {
  return Urql.useQuery<GetUsersQuery, GetUsersQueryVariables>({ query: GetUsersDocument, ...options });
};
export const GetUserDocument = gql`
    query GetUser($userId: Int!) {
  getUser(userId: $userId) {
    user {
      username
      userId
      role
    }
  }
}
    `;

export function useGetUserQuery(options: Omit<Urql.UseQueryArgs<GetUserQueryVariables>, 'query'>) {
  return Urql.useQuery<GetUserQuery, GetUserQueryVariables>({ query: GetUserDocument, ...options });
};
export const GetUserByUsernameDocument = gql`
    query GetUserByUsername($username: String!) {
  getUserByUsername(username: $username) {
    user {
      username
      userId
      role
    }
  }
}
    `;

export function useGetUserByUsernameQuery(options: Omit<Urql.UseQueryArgs<GetUserByUsernameQueryVariables>, 'query'>) {
  return Urql.useQuery<GetUserByUsernameQuery, GetUserByUsernameQueryVariables>({ query: GetUserByUsernameDocument, ...options });
};
export const RegisterUserDocument = gql`
    mutation RegisterUser($input: RegisterUserInput!) {
  registerUser(input: $input) {
    message
    success
  }
}
    `;

export function useRegisterUserMutation() {
  return Urql.useMutation<RegisterUserMutation, RegisterUserMutationVariables>(RegisterUserDocument);
};
export const UpdateUserDocument = gql`
    mutation UpdateUser($input: UpdateUserInput!) {
  updateUser(input: $input) {
    message
    success
  }
}
    `;

export function useUpdateUserMutation() {
  return Urql.useMutation<UpdateUserMutation, UpdateUserMutationVariables>(UpdateUserDocument);
};
export const DeleteUserDocument = gql`
    mutation DeleteUser($userId: Int!) {
  deleteUser(userId: $userId) {
    message
    success
  }
}
    `;

export function useDeleteUserMutation() {
  return Urql.useMutation<DeleteUserMutation, DeleteUserMutationVariables>(DeleteUserDocument);
};