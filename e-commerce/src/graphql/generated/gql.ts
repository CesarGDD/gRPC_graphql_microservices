/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "query GetOrdersDetailsByUserId($input: GetOrdersDetailsByUserIdInput!) {\n  getOrdersDetailsByUserId(input: $input) {\n    orderId\n    orderItems {\n      productId\n      quantity\n    }\n    status\n    totalPrice\n    userId\n  }\n}\n\nquery getOrderDetails($input: GetOrderDetailsInput!) {\n  getOrderDetails(input: $input) {\n    orderDetails {\n      orderId\n      orderItems {\n        productId\n        quantity\n      }\n      status\n      totalPrice\n      userId\n    }\n  }\n}\n\nquery getPaymentDetails($input: GetPaymentDetailsInput!) {\n  getPaymentDetails(input: $input) {\n    paymentMethod\n  }\n}\n\nmutation createOrder($input: CreateOrderInput!) {\n  createOrder(input: $input) {\n    message\n    success\n  }\n}\n\nmutation processPayment($input: ProcessPaymentInput!) {\n  processPayment(input: $input) {\n    message\n    success\n  }\n}": types.GetOrdersDetailsByUserIdDocument,
    "query GetProducts {\n  getProducts {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nquery getProductByName($name: String!) {\n  getProductByName(name: $name) {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nquery GetProduct($productId: Int!) {\n  getProduct(productId: $productId) {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nmutation CreateProduct($input: NewProductInput!) {\n  createProduct(input: $input) {\n    message\n    success\n  }\n}\n\nmutation DeleteProduct($productId: Int!) {\n  deleteProduct(productId: $productId) {\n    message\n    success\n  }\n}\n\nmutation UpdateProduct($input: UpdateProductInput!) {\n  updateProduct(input: $input) {\n    message\n    success\n  }\n}": types.GetProductsDocument,
    "query GetCart($input: GetCartInput!) {\n  getCart(input: $input) {\n    items {\n      product {\n        description\n        name\n        price\n        productId\n        title\n        url\n        userId\n      }\n      productId\n      quantity\n    }\n    userId\n  }\n}\n\nmutation AddItem($input: AddItemInput!) {\n  addItem(input: $input) {\n    message\n    success\n  }\n}\n\nmutation ClearCart($input: ClearCartInput!) {\n  clearCart(input: $input) {\n    message\n    success\n  }\n}\n\nmutation RemoveItem($input: RemoveItemInput!) {\n  removeItem(input: $input) {\n    message\n    success\n  }\n}\n\nmutation UpdateItem($input: UpdateItemInput!) {\n  updateItem(input: $input) {\n    message\n    success\n  }\n}": types.GetCartDocument,
    "query GetUsers {\n  getUsers {\n    username\n    userId\n    role\n  }\n}\n\nquery GetUser($userId: Int!) {\n  getUser(userId: $userId) {\n    username\n    userId\n    role\n  }\n}\n\nquery GetUserByUsername($username: String!) {\n  getUserByUsername(username: $username) {\n    username\n    userId\n    role\n  }\n}\n\nmutation Login($input: LoginInput!) {\n  login(input: $input) {\n    token\n    user {\n      userId\n      username\n    }\n  }\n}\n\nmutation SignIn($input: RegisterUserInput) {\n  signIn(input: $input) {\n    token\n    user {\n      userId\n      username\n    }\n  }\n}\n\nmutation UpdateUser($input: UpdateUserInput!) {\n  updateUser(input: $input) {\n    message\n    success\n  }\n}\n\nmutation DeleteUser($userId: Int!) {\n  deleteUser(userId: $userId) {\n    message\n    success\n  }\n}": types.GetUsersDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query GetOrdersDetailsByUserId($input: GetOrdersDetailsByUserIdInput!) {\n  getOrdersDetailsByUserId(input: $input) {\n    orderId\n    orderItems {\n      productId\n      quantity\n    }\n    status\n    totalPrice\n    userId\n  }\n}\n\nquery getOrderDetails($input: GetOrderDetailsInput!) {\n  getOrderDetails(input: $input) {\n    orderDetails {\n      orderId\n      orderItems {\n        productId\n        quantity\n      }\n      status\n      totalPrice\n      userId\n    }\n  }\n}\n\nquery getPaymentDetails($input: GetPaymentDetailsInput!) {\n  getPaymentDetails(input: $input) {\n    paymentMethod\n  }\n}\n\nmutation createOrder($input: CreateOrderInput!) {\n  createOrder(input: $input) {\n    message\n    success\n  }\n}\n\nmutation processPayment($input: ProcessPaymentInput!) {\n  processPayment(input: $input) {\n    message\n    success\n  }\n}"): (typeof documents)["query GetOrdersDetailsByUserId($input: GetOrdersDetailsByUserIdInput!) {\n  getOrdersDetailsByUserId(input: $input) {\n    orderId\n    orderItems {\n      productId\n      quantity\n    }\n    status\n    totalPrice\n    userId\n  }\n}\n\nquery getOrderDetails($input: GetOrderDetailsInput!) {\n  getOrderDetails(input: $input) {\n    orderDetails {\n      orderId\n      orderItems {\n        productId\n        quantity\n      }\n      status\n      totalPrice\n      userId\n    }\n  }\n}\n\nquery getPaymentDetails($input: GetPaymentDetailsInput!) {\n  getPaymentDetails(input: $input) {\n    paymentMethod\n  }\n}\n\nmutation createOrder($input: CreateOrderInput!) {\n  createOrder(input: $input) {\n    message\n    success\n  }\n}\n\nmutation processPayment($input: ProcessPaymentInput!) {\n  processPayment(input: $input) {\n    message\n    success\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query GetProducts {\n  getProducts {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nquery getProductByName($name: String!) {\n  getProductByName(name: $name) {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nquery GetProduct($productId: Int!) {\n  getProduct(productId: $productId) {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nmutation CreateProduct($input: NewProductInput!) {\n  createProduct(input: $input) {\n    message\n    success\n  }\n}\n\nmutation DeleteProduct($productId: Int!) {\n  deleteProduct(productId: $productId) {\n    message\n    success\n  }\n}\n\nmutation UpdateProduct($input: UpdateProductInput!) {\n  updateProduct(input: $input) {\n    message\n    success\n  }\n}"): (typeof documents)["query GetProducts {\n  getProducts {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nquery getProductByName($name: String!) {\n  getProductByName(name: $name) {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nquery GetProduct($productId: Int!) {\n  getProduct(productId: $productId) {\n    description\n    name\n    price\n    productId\n    title\n    url\n    userId\n    userOwner {\n      userId\n      username\n    }\n  }\n}\n\nmutation CreateProduct($input: NewProductInput!) {\n  createProduct(input: $input) {\n    message\n    success\n  }\n}\n\nmutation DeleteProduct($productId: Int!) {\n  deleteProduct(productId: $productId) {\n    message\n    success\n  }\n}\n\nmutation UpdateProduct($input: UpdateProductInput!) {\n  updateProduct(input: $input) {\n    message\n    success\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query GetCart($input: GetCartInput!) {\n  getCart(input: $input) {\n    items {\n      product {\n        description\n        name\n        price\n        productId\n        title\n        url\n        userId\n      }\n      productId\n      quantity\n    }\n    userId\n  }\n}\n\nmutation AddItem($input: AddItemInput!) {\n  addItem(input: $input) {\n    message\n    success\n  }\n}\n\nmutation ClearCart($input: ClearCartInput!) {\n  clearCart(input: $input) {\n    message\n    success\n  }\n}\n\nmutation RemoveItem($input: RemoveItemInput!) {\n  removeItem(input: $input) {\n    message\n    success\n  }\n}\n\nmutation UpdateItem($input: UpdateItemInput!) {\n  updateItem(input: $input) {\n    message\n    success\n  }\n}"): (typeof documents)["query GetCart($input: GetCartInput!) {\n  getCart(input: $input) {\n    items {\n      product {\n        description\n        name\n        price\n        productId\n        title\n        url\n        userId\n      }\n      productId\n      quantity\n    }\n    userId\n  }\n}\n\nmutation AddItem($input: AddItemInput!) {\n  addItem(input: $input) {\n    message\n    success\n  }\n}\n\nmutation ClearCart($input: ClearCartInput!) {\n  clearCart(input: $input) {\n    message\n    success\n  }\n}\n\nmutation RemoveItem($input: RemoveItemInput!) {\n  removeItem(input: $input) {\n    message\n    success\n  }\n}\n\nmutation UpdateItem($input: UpdateItemInput!) {\n  updateItem(input: $input) {\n    message\n    success\n  }\n}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query GetUsers {\n  getUsers {\n    username\n    userId\n    role\n  }\n}\n\nquery GetUser($userId: Int!) {\n  getUser(userId: $userId) {\n    username\n    userId\n    role\n  }\n}\n\nquery GetUserByUsername($username: String!) {\n  getUserByUsername(username: $username) {\n    username\n    userId\n    role\n  }\n}\n\nmutation Login($input: LoginInput!) {\n  login(input: $input) {\n    token\n    user {\n      userId\n      username\n    }\n  }\n}\n\nmutation SignIn($input: RegisterUserInput) {\n  signIn(input: $input) {\n    token\n    user {\n      userId\n      username\n    }\n  }\n}\n\nmutation UpdateUser($input: UpdateUserInput!) {\n  updateUser(input: $input) {\n    message\n    success\n  }\n}\n\nmutation DeleteUser($userId: Int!) {\n  deleteUser(userId: $userId) {\n    message\n    success\n  }\n}"): (typeof documents)["query GetUsers {\n  getUsers {\n    username\n    userId\n    role\n  }\n}\n\nquery GetUser($userId: Int!) {\n  getUser(userId: $userId) {\n    username\n    userId\n    role\n  }\n}\n\nquery GetUserByUsername($username: String!) {\n  getUserByUsername(username: $username) {\n    username\n    userId\n    role\n  }\n}\n\nmutation Login($input: LoginInput!) {\n  login(input: $input) {\n    token\n    user {\n      userId\n      username\n    }\n  }\n}\n\nmutation SignIn($input: RegisterUserInput) {\n  signIn(input: $input) {\n    token\n    user {\n      userId\n      username\n    }\n  }\n}\n\nmutation UpdateUser($input: UpdateUserInput!) {\n  updateUser(input: $input) {\n    message\n    success\n  }\n}\n\nmutation DeleteUser($userId: Int!) {\n  deleteUser(userId: $userId) {\n    message\n    success\n  }\n}"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;