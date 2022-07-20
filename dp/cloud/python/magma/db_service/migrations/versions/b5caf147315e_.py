"""empty message

Revision ID: b5caf147315e
Revises: 00d77c9f7532
Create Date: 2022-06-20 15:44:14.272618

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = 'b5caf147315e'
down_revision = '00d77c9f7532'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('cbsds', sa.Column('carrier_aggregation_enabled', sa.Boolean(), server_default='false', nullable=False))
    op.add_column('cbsds', sa.Column('max_ibw_mhz', sa.Integer(), server_default='150', nullable=False))
    op.add_column('cbsds', sa.Column('grant_redundancy', sa.Boolean(), server_default='true', nullable=False))
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_column('cbsds', 'grant_redundancy')
    op.drop_column('cbsds', 'max_ibw_mhz')
    op.drop_column('cbsds', 'carrier_aggregation_enabled')
    # ### end Alembic commands ###
