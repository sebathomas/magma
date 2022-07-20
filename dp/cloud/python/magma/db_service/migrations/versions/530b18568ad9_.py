"""empty message

Revision ID: 530b18568ad9
Revises: b5caf147315e
Create Date: 2022-06-27 15:55:48.080182

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = '530b18568ad9'
down_revision = 'b5caf147315e'
branch_labels = None
depends_on = None


def upgrade():
    """
    Run upgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('cbsds', sa.Column('available_frequencies', sa.JSON(), nullable=True))
    # ### end Alembic commands ###


def downgrade():
    """
    Run downgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_column('cbsds', 'available_frequencies')
    # ### end Alembic commands ###
