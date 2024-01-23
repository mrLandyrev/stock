create table placements (
	stock_id UUID not null,
	product_id UUID not null,
	count int not null,
	reserved int not null,
	primary key(stock_id, product_id),
	constraint fk_stock foreign key(stock_id) references stocks(id),
	constraint fk_product foreign key(product_id) references products(id)
)