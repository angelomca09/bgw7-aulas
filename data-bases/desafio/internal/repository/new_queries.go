package repository

const (
	UpdateInvoiceTotals = `
    UPDATE invoices i
    SET i.total = (
        SELECT COALESCE(SUM(p.price * s.quantity), 0)
        FROM sales s
            INNER JOIN products p ON p.id = s.product_id
        WHERE s.invoice_id = i.id
    );
    `

	GetTotalByCondition = `
    SELECT
    c.condition, SUM(i.total)
    FROM customers c
        INNER JOIN invoices i ON c.id = i.customer_id
    GROUP BY c.condition;
    `

	GetTop5SoldProducts = `
    SELECT
    p.description, SUM(s.quantity) AS total_sold
    FROM products p
        INNER JOIN sales s ON p.id = s.product_id
    GROUP BY p.id
    ORDER BY total_sold DESC
    LIMIT 5;
    `

	GetTop5ActiveCustomersByTotalSpent = `
    SELECT
    c.first_name, c.last_name, SUM(i.total) AS amount
    FROM customers c
        INNER JOIN invoices i ON c.id = i.customer_id
    GROUP BY c.id
    ORDER BY amount DESC
    LIMIT 5;
    `
)
