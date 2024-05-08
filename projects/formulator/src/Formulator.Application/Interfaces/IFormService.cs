using Formulator.Core.Entities;

namespace Formulator.Application.Interfaces
{
    public interface IFormService
    {
        Form CreateForm(Form form);

        Form GetForm(int id);
        IEnumerable<Form> GetUserForms(User user);

        void UpdateForm(Form form);
        void DeleteForm(int id);
    }
}