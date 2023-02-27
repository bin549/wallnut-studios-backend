import http
from flask_restx import Namespace,Resource,fields
from flask_jwt_extended import jwt_required,get_jwt_identity
from ..models.orders import Order
from ..models.users import User
from http import HTTPStatus
from ..utils.db import db


order_namespace=Namespace('orders',description="Namespace for orders")

order_model=order_namespace.model(
    'Order',{
        'id':fields.Integer(description="An ID"),
        'size':fields.String(description="Size of order",required=True,
            enum=['SMALL','MEDIUM','LARGE','EXTRA_LARGE']
        ),
        'order_status':fields.String(description="The status of the Order",
            required=True, enum=['PENDING','IN_TRANSIT','DELIVERED']
        )
    }
)



@order_namespace.route('/orders/')
class OrderGetCreate(Resource):
    @order_namespace.marshal_with(order_model)
    @order_namespace.doc(description="Retrieve all orders")
    @jwt_required()
    def get(self):
        orders=Order.query.all()
        return orders, HTTPStatus.OK

    @order_namespace.expect(order_model)
    @order_namespace.marshal_with(order_model)
    @order_namespace.doc(
        description="Place an order"
    )
    @jwt_required()
    def post(self):
        data=order_namespace.payload
        username=get_jwt_identity()
        print(username)
        current_user=User.query.filter_by(username=username).first()
        new_order=Order(
            size=data['size'],
            quantity=data['quantity'],
            flavour=data['flavour']
        )
        new_order.user=current_user
        new_order.save()
        return new_order , HTTPStatus.CREATED


@order_namespace.route('/order/<int:order_id>')
class GetUpdateDelete(Resource):
    def get(self):
        pass

    def put(self):
        pass

    def delete(self):
        pass


@order_namespace.route('/user/<int:user_id>/order/<int:order_id>/')
class GetSpecificOrderByUser(Resource):
    def get(self):
        pass


@order_namespace.route('/order/status/<int:order_id>')
class UpdateOrderStatus(Resource):
    def patch(self, order_id):
        pass
