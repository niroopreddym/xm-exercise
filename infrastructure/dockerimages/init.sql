CREATE TABLE public.Company (
    id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name varchar(255),
    code varchar(255),
    country varchar(255),
    website varchar(255),
    phone varchar(255)
);

-- CREATE OR REPLACE FUNCTION AddProductsQuantity (productid int,quantity int)
-- returns void AS
-- $body$ 
-- DECLARE
-- 	existingquantity int;
-- BEGIN
--   	 existingquantity :=(select availablequantity from public.products where id = productid);
--      update public.products set availablequantity =  quantity + existingquantity  where id = productid;
-- END;
-- $body$
-- LANGUAGE plpgsql;

-- CREATE OR REPLACE FUNCTION BookProducts (productid int,quantity int)
-- returns boolean AS
-- $body$ 
-- DECLARE
-- 	existingquantity int;
-- BEGIN
--   	 existingquantity :=(select availablequantity from public.products where id = productid);
-- 	 if existingquantity < quantity then
-- 	 	return false;
-- 	else
--     	update public.products set availablequantity =  existingquantity-quantity  where id = productid;
-- 		return true;
-- 	end if;
-- END;
-- $body$
-- LANGUAGE plpgsql;