import React from "react";
import { useDispatch, useSelector } from "react-redux";

import Grid from "@material-ui/core/Grid";

import CallToAction from "../components/CallToAction";
import Products from "../components/Products";

import { setDialog } from "../state/features/dialog";
import { useGetAllProductsQuery } from "../state/service";

const products = [
  {
    id: 0,
    name: "Xiaomi Redmi Note 4",
    user: { id: 0, name: "Khaled Emara" },
    price: 40.0,
    image_url:
      "https://fdn2.gsmarena.com/vv/pics/xiaomi/xiaomi-redmi-note4-sn-1.jpg",
    description: "A budget phone.",
  },
  {
    id: 1,
    name: "Samsung EarBuds+",
    user: { id: 1, name: "Hosni Adel" },
    price: 30.0,
    image_url:
      "https://images.samsung.com/is/image/samsung/uk/galaxy-s20/gallery/uk-galaxy-buds-plus-sm-r175nzkaeua-casetopcombinationblack-208767354?$2052_1641_PNG$",
    description: "An in-ear headset.",
  },
  {
    id: 2,
    name: "Air Purifier",
    user: { id: 2, name: "Hussein Moustafa" },
    price: 50.0,
    image_url:
      "https://www.lg.com/eg_en/images/air-purifiers/md07515500/gallery/AS95GDWV0-L1.jpg",
    description: "Filters PM2.5 particles and harmful gases.",
  },
];

function Home() {
  const dispatch = useDispatch();
  const current_user = useSelector((state) => state.user);

  const { data, error, isLoading } = useGetAllProductsQuery();

  const handleAction = () => {
    if (current_user.id !== null) {
      dispatch(setDialog("add-product"));
    } else {
      dispatch(setDialog("sign-in"));
    }
  };

  return (
    <Grid container spacing={4}>
      <Grid item xs={12}>
        <CallToAction
          headerTitle="Distributed Marketplace"
          subheaderTitle="Buy and sell on a highly scalable marketplace"
          subtitle={
            current_user.id !== null
              ? `Start selling now`
              : "Sign in and start selling now"
          }
          primaryActionText={current_user.id !== null ? "Sell" : "Sign in"}
          handlePrimaryAction={handleAction}
        />
      </Grid>
      <Grid item xs={12}>
        {isLoading ? (
          <div>Loading...</div>
        ) : error ? (
          <div>
            {error.data && error.data.error ? error.data.error : error.status}
          </div>
        ) : data && data.data && data.data.length > 0 ? (
          <Products products={data.data} />
        ) : (
          <Products products={products} />
        )}
      </Grid>
    </Grid>
  );
}

export default Home;
