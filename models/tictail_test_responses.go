package models

const (
	TICTAIL_MOCK_GET_ME_200_RESPONSE = `{
  "currency": "SEK",
  "id": "3MpE",
  "dashboard_url": "https://tictail.com/dashboard/store/testuser",
  "logotype": null,
  "state": null,
  "followers": 2,
  "storekeeper_email": "test@testuser.com",
  "online": true,
  "short_description": "",
  "subdomain": "testuser",
  "launched_at": "2015-08-19T08:53:41.000000",
  "vat": 
    "applied_to_shipping": false,
    "rate": "0.060000",
    "region": "SE",
    "included_in_prices": true
  },
  "wallpapers": {
    "iphone": {
      "original_height": 640,
      "sizes": {
        "640": "https://images.ttcdn.co/media/i/wallpaper/311676-5aef267edab14a7da8bf470d0e0d1d17.jpeg?size=640"
      },
      "url": "https://images.ttcdn.co/media/i/wallpaper/311676-5aef267edab14a7da8bf470d0e0d1d17.jpeg",
      "created_at": "2015-08-09T12:35:05.000000",
      "modified_at": "2015-08-09T12:35:05.000000",
      "crop_width": 640,
      "id": "d3c",
      "original_width": 640,
      "crop_height": 640,
      "crop_y": 0,
      "crop_x": 0
    }
  },
  "description": "",
 "contact_email": "test@testuser.com,
  "name": "Test Store",
  "language": "sv",
  "url": "http://testuser.tictail.com",
  "country": "SE",
  "created_at": "2015-07-19T06:51:45.000000",
  "modified_at": null,
  "appstore_currency": "USD",
  "sandbox": false
}`
	TICTAIL_MOCK_GET_PRODUCTS_200_RESPONSE = `[{
    "status": "published",
    "store_url": "http://testuser.tictail.com",
    "description": "",
    "store_name": "TestUser's Store",
    "store_id": "3MpE",
    "unlimited": true,
    "created_at": "2015-09-16T15:07:05.000000",
    "title": "TestProduct",
    "modified_at": "2015-09-23T18:44:12.000000",
    "slug": "testproduct",
    "price": 9900,
    "currency": "SEK",
    "store_subdomain": "testuser",
    "likes": 0,
    "variations": [
      {
        "title": null,
        "modified_at": "2015-09-23T09:53:53.000000",
        "created_at": "2015-09-16T15:07:05.000000",
        "unlimited": true,
        "id": "Ek2Y",
        "quantity": null
      }
    ],
    "images": [
      {
        "original_height": 2048,
        "sizes": {
          "2000": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=2000",
          "30": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=30",
          "300": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=300",
          "45": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=45",
          "50": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=50",
          "40": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=40",
          "640": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=640",
          "75": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=75",
          "100": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=100",
          "500": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=500",
          "1000": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png?size=1000"
        },
        "url": "https://images.ttcdn.co/media/i/product/311676-6543a7898516436daa620e362620f643.png",
        "created_at": "2015-09-20T12:55:45.000000",
        "modified_at": "2015-09-20T12:55:45.000000",
        "original_width": 2048,
        "id": "MjzN"
      },
      {
        "original_height": 2048,
        "sizes": {
          "2000": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=2000",
          "30": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=30",
          "300": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=300",
          "45": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=45",
          "50": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=50",
          "40": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=40",
          "640": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=640",
          "75": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=75",
          "100": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=100",
          "500": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=500",
          "1000": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png?size=1000"
        },
        "url": "https://images.ttcdn.co/media/i/product/311676-f9346272ee8545e499934345c960cf2f.png",
        "created_at": "2015-09-20T12:55:45.000000",
        "modified_at": "2015-09-20T12:55:45.000000",
        "original_width": 2048,
        "id": "MjzP"
      },
      {
        "original_height": 2048,
        "sizes": {
          "2000": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=2000",
          "30": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=30",
          "300": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=300",
          "45": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=45",
          "50": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=50",
          "40": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=40",
          "640": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=640",
          "75": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=75",
          "100": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=100",
          "500": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=500",
          "1000": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png?size=1000"
        },
        "url": "https://images.ttcdn.co/media/i/product/311676-1eb2d98f09b5481ea35723a42e9e3245.png",
        "created_at": "2015-09-20T12:55:45.000000",
        "modified_at": "2015-09-20T12:55:45.000000",
        "original_width": 2048,
        "id": "MjzQ"
      }
    ],
    "id": "mbYi",
    "categories": [
      {
        "title": "TestProduct",
        "created_at": "2015-09-20T12:55:37.000000",
        "modified_at": null,
        "slug": "testproduct",
        "parent_id": null,
        "product_count": 1,
        "position": 0,
        "id": "6PCE"
      }
    ],
    "quantity": null
  }
]`
	TICTAIL_MOCK_GET_ORDERS_200_RESPONSE = `[{
    "customer": {
      "name": "Test Customer",
      "language": "en",
      "country": "SE",
      "created_at": "2015-09-25T10:04:38.000000",
      "modified_at": null,
      "id": "cLnj",
      "email": "test.test@gmail.com"
    },
    "transaction": {
      "reference": "xxx",
      "status": "paid",
      "paid_at": "2015-09-25T10:06:40.680296",
      "pending_reason": null,
      "processor": "stripe"
    },
    "attribution": "storefront",
    "modified_at": "2015-09-25T10:06:44.083114",
    "discounts": [],
    "items": [
      {
        "currency": "SEK",
        "price": 9900,
        "product": {
          "status": "published",
          "store_url": null,
          "description": ""
          "store_name": null,
          "store_id": "3MpE",
          "unlimited": true,
          "created_at": "2015-09-16T15:07:05.000000",
          "title": "TestProduct",
          "modified_at": "2015-09-23T18:44:12.000000",
          "variation": null,
          "slug": "testproduct",
          "price": 9900,
          "currency": "SEK",
          "store_subdomain": null,
          "likes": 0,
          "images": [
          ],
          "id": "mbYi",
          "categories": [
            {
              "title": "Product Title",
              "created_at": "2015-09-20T12:55:37.000000",
              "modified_at": null,
              "slug": "product",
              "parent_id": null,
              "product_count": 1,
              "position": 0,
              "id": "6PCE"
            }
          ],
          "quantity": null
        },
        "quantity": 1
      }
    ],
    "completed": "false",
    "invoice_fee": 0,
    "messages": [],
    "number": 4109746,
    "id": "rqwk",
    "note": null,
    "currency": "SEK",
    "shipping_alternative": {
      "maximum_delivery_days": 4,
      "regions": [
        "SE"
      ],
      "excluded_regions": [],
      "description": null,
      "free_at_price": null,
      "created_at": "2015-07-30T23:25:32.000000",
      "title": "Frakt",
      "modified_at": "2015-09-23T18:35:10.000000",
      "requires_customer_phone": false,
      "price": 4900,
      "currency": "SEK",
      "id": "3paE",
      "minimum_delivery_days": 2
    },
    "fullfilment": {
      "status": "unhandled",
      "price": 0,
      "provider": null,
      "currency": "SEK",
      "tracking_number": null,
      "receiver": {
        "city": "test",
        "street_line2": null,
        "name": "Test Customer",
        "zip": "",
        "country": "SE",
        "phone": null,
        "state": null,
        "street": ""
      },
      "shipped_at": null,
      "vat": {
        "currency": "SEK",
        "price": 0,
        "rate": "0.000000"
      }
    },
    "price": 9900,
    "created_at": "2015-09-25T10:06:40.639978",
    "vat": {
      "currency": "SEK",
      "price": 560,
      "rate": "0.059957"
    },
    "prices_include_vat": true
  }
]`
)
