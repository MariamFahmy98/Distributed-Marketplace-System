import React from "react";
import { useParams } from "react-router-dom";

import Container from "@material-ui/core/Container";

import { Product } from "../components/Products";

import { useGetProductByIdQuery } from "../state/service";

function ProductPage() {
  const params = useParams();

  const { data, error, isLoading } = useGetProductByIdQuery(params.id);

  return (
    <Container maxWidth="sm">
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>
          {error.data && error.data.error ? error.data.error : error.status}
        </div>
      ) : data && data.data ? (
        <Product {...data.data} />
      ) : (
        <div>Couldn't display product</div>
      )}
    </Container>
  );
}

export default ProductPage;
