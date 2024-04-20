usermanagment-svc:
	cp usermanagement.proto ../services/usermanagement/proto/ && \
	cp buf.gen.yaml ../services/usermanagement/proto/ && \
	cd ../services/usermanagement/proto && \
	buf generate

products-svc:
	cp products.proto ../services/products/proto/ && \
	cp buf.gen.yaml ../services/products/proto/ && \
	cd ../services/products/proto && \
	buf generate

shoppingcart-svc:
	cp shoppingcart.proto ../services/shoppingcart/proto/ && \
	cp buf.gen.yaml ../services/shoppingcart/proto/ && \
	cd ../services/shoppingcart/proto && \
	buf generate

checkout-svc:
	cp checkout.proto ../services/checkout/proto/ && \
	cp buf.gen.yaml ../services/checkout/proto/ && \
	cd ../services/checkout/proto && \
	buf generate

graphql_api:
	cp checkout.proto ../graphql_api/protos/checkoutpb && \
	cp shoppingcart.proto ../graphql_api/protos/shoppingcartpb && \
	cp products.proto ../graphql_api/protos/productspb && \
	cp usermanagement.proto ../graphql_api/protos/usermanagementpb && \
	cp buf.gen.yaml ../graphql_api/protos/ && \
	cd ../graphql_api/protos && \
	buf generate
