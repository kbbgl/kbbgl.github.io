# Filtering Nested Objects

TODO clean up

Let's say we have the following paths with interesting information:

```javascript
${Marketplace.SubscriptionService.Transaction.xrefCData.vendorId}
${Marketplace.SubscriptionService.Transaction.xrefCData.vendorName}
${Marketplace.SubscriptionService.Transaction.xrefCData.transactionType}
${Marketplace.SubscriptionService.Transaction.xrefPackId}
```

If we want to extract the nested objects based on the filter `vendor: 123`:

```javascript
!Print value="${PartnerReportTransactions(val.xrefCData.vendorId == '123')}" extend-context=`TransactionsSummary=.=[{"transactiontype": val.xrefCData.transactionType, "pack": val.xrefPackId, "transactionid": val.xrefTransactionId}]` ignore-outputs=true
```

## Initializing a Grid

We can then initialize

```javascript
!setIncident marketplacetransactionstable=[]
```

```javascript
!SetGridField columns="Transaction Type,Pack,Transaction ID" grid_id="Marketplace Transactions Table" context_path=TransactionsSummary overwrite=true keys="transactiontype,pack,transactionid"
```

## Create new Context with filter

```javascript

// Get report
!mp-subscriptions-get-report month=${MONTH} year={YEAR} using="Marketplace Susbcription Production"

// Get all vendors
!Print value=${Marketplace.SubscriptionService.Transaction.xrefCData.vendorId}

// Create keys per vendor
!Set key=v_txs value="${Marketplace.SubscriptionService.Transaction(val.xrefCData.vendorId == $VENDOR_ID)}"
!MarketplacePartnerMonthlyReport partner_name="v" transactions=${v_txs} month=August year=2022
```
